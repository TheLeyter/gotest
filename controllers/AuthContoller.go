package controllers

import (
	"gotest/services/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiController interface {
	InitApiRoutes()
}

type AuthContoller struct {
	Routes  *gin.RouterGroup
	Service *auth.AuthService
}

func NewAuthController(Routes *gin.RouterGroup, Service *auth.AuthService) ApiController {
	controller := AuthContoller{Routes: Routes, Service: Service}
	return &controller
}

func (controller *AuthContoller) InitApiRoutes() {
	authGroup := controller.Routes.Group("/auth")

	authGroup.POST("/login", controller.login)
	authGroup.POST("/register", controller.register)

}

func (controller *AuthContoller) login(c *gin.Context) {
	var loginRequest auth.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	response, err := controller.Service.Login(&loginRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (controller *AuthContoller) register(c *gin.Context) {
	var registerRequest auth.RegisterRequest

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	response, err := controller.Service.Register(&registerRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
