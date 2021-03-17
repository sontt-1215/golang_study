package controller

import (
    // "net/http"
    // "strconv"
    "fmt"

    "github.com/gin-gonic/gin"
    // "golang_api/app/http/request"
    "golang_api/app/service"
)

type Office365Controller interface {
    StartAuth(ctx *gin.Context)
    CallbackAuth(ctx *gin.Context)
}

type office365Controller struct {
    office365Service service.Office365Service
    jwtService  service.JWTService
}

func NewOffice365Controller(office365Service service.Office365Service, jwtService service.JWTService) Office365Controller {
    return &office365Controller{
        office365Service: office365Service,
        jwtService: jwtService,
    }
}

func (o *office365Controller) StartAuth(ctx *gin.Context) {
    fmt.Println("aloooo")
    // var loginRequest request.LoginRequest
    // errRequest := ctx.ShouldBind(&loginRequest)
    // if errRequest != nil {
    //     response := helper.BuildErrorResponse("Failed to process request", errRequest.Error(), helper.EmptyObj{})
    //     ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
    //     return
    // }
    // authResult := c.authService.VerifyCredential(loginRequest.Email, loginRequest.Password)
    // if v, ok := authResult.(model.User); ok {
    //     generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
    //     v.Token = generatedToken
    //     response := helper.BuildResponse(true, "OK!", v)
    //     ctx.JSON(http.StatusOK, response)
    //     return
    // }
    // response := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{})
    // ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (o *office365Controller) CallbackAuth(ctx *gin.Context) {
    // var registerRequest request.RegisterRequest
    // errRequest := ctx.ShouldBind(&registerRequest)
    // if errRequest != nil {
    //     response := helper.BuildErrorResponse("Failed to process request", errRequest.Error(), helper.EmptyObj{})
    //     ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
    //     return
    // }

    // if !c.authService.IsDuplicateEmail(registerRequest.Email) {
    //     response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
    //     ctx.JSON(http.StatusConflict, response)
    // } else {
    //     createdUser := c.authService.CreateUser(registerRequest)
    //     token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
    //     createdUser.Token = token
    //     response := helper.BuildResponse(true, "OK!", createdUser)
    //     ctx.JSON(http.StatusCreated, response)
    // }
}
