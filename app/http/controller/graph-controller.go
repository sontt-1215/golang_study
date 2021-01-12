package controller

import (
    "fmt"
    "net/http"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "golang_api/helper"
    "golang_api/app/service"
)

type GraphController interface {
    calendar(context *gin.Context)
}

type graphController struct {
    graphService service.GraphService
    jwtService  service.JWTService
}

func NewGraphController(graphService service.GraphService, jwtService service.JWTService) GraphController {
    return &graphController{
        graphService: graphService,
        jwtService:  jwtService,
    }
}

func (c *graphController) Calendar(context *gin.Context) {
    authHeader := context.GetHeader("Authorization")
    token, err := c.jwtService.ValidateToken(authHeader)
    if err != nil {
        panic(err.Error())
    }
    claims := token.Claims.(jwt.MapClaims)
    id := fmt.Sprintf("%v", claims["user_id"])
    graph := c.graphService.Profile(id)
    res := helper.BuildResponse(true, "OK", graph)
    context.JSON(http.StatusOK, res)

}
