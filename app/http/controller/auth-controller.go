package controller

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "golang_api/app/http/request"
    "golang_api/app/model"
    "golang_api/helper"
    "golang_api/app/service"
)

type AuthController interface {
    Login(ctx *gin.Context)
    Register(ctx *gin.Context)
}

type authController struct {
    authService service.AuthService
    jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
    return &authController{
        authService: authService,
        jwtService:  jwtService,
    }
}

func (c *authController) Login(ctx *gin.Context) {
    var loginRequest request.LoginRequest
    errRequest := ctx.ShouldBind(&loginRequest)
    if errRequest != nil {
        response := helper.BuildErrorResponse("Failed to process request", errRequest.Error(), helper.EmptyObj{})
        ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
        return
    }
    authResult := c.authService.VerifyCredential(loginRequest.Email, loginRequest.Password)
    if v, ok := authResult.(model.User); ok {
        generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
        v.Token = generatedToken
        response := helper.BuildResponse(true, "OK!", v)
        ctx.JSON(http.StatusOK, response)
        return
    }
    response := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{})
    ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
    var registerRequest request.RegisterRequest
    errRequest := ctx.ShouldBind(&registerRequest)
    if errRequest != nil {
        response := helper.BuildErrorResponse("Failed to process request", errRequest.Error(), helper.EmptyObj{})
        ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
        return
    }

    if !c.authService.IsDuplicateEmail(registerRequest.Email) {
        response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
        ctx.JSON(http.StatusConflict, response)
    } else {
        createdUser := c.authService.CreateUser(registerRequest)
        token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
        createdUser.Token = token
        response := helper.BuildResponse(true, "OK!", createdUser)
        ctx.JSON(http.StatusCreated, response)
    }
}
