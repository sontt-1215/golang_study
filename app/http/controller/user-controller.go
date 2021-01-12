package controller

import (
    "fmt"
    "net/http"
    "strconv"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "golang_api/app/http/request"
    "golang_api/helper"
    "golang_api/app/service"
)

type UserController interface {
    Update(context *gin.Context)
    Profile(context *gin.Context)
}

type userController struct {
    userService service.UserService
    jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
    return &userController{
        userService: userService,
        jwtService:  jwtService,
    }
}

func (c *userController) Update(context *gin.Context) {
    var userUpdateRequest request.UserUpdateRequest
    errRequest := context.ShouldBind(&userUpdateRequest)
    if errRequest != nil {
        res := helper.BuildErrorResponse("Failed to process request", errRequest.Error(), helper.EmptyObj{})
        context.AbortWithStatusJSON(http.StatusBadRequest, res)
        return
    }

    authHeader := context.GetHeader("Authorization")
    token, errToken := c.jwtService.ValidateToken(authHeader)
    if errToken != nil {
        panic(errToken.Error())
    }
    claims := token.Claims.(jwt.MapClaims)
    id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
    if err != nil {
        panic(err.Error())
    }
    userUpdateRequest.ID = id
    u := c.userService.Update(userUpdateRequest)
    res := helper.BuildResponse(true, "OK!", u)
    context.JSON(http.StatusOK, res)
}

func (c *userController) Profile(context *gin.Context) {
    authHeader := context.GetHeader("Authorization")
    token, err := c.jwtService.ValidateToken(authHeader)
    if err != nil {
        panic(err.Error())
    }
    claims := token.Claims.(jwt.MapClaims)
    id := fmt.Sprintf("%v", claims["user_id"])
    user := c.userService.Profile(id)
    res := helper.BuildResponse(true, "OK", user)
    context.JSON(http.StatusOK, res)

}
