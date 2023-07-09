// controllers/home_controller.go
package controllers

import (
	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func (hc *HomeController) Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Добро пожаловать на домашнюю страницу!",
	})
}

// controllers/profile_controller.go
package controllers

import (
	"github.com/gin-gonic/gin"
)

type ProfileController struct{}

func (pc *ProfileController) GetProfile(c *gin.Context) {
	username := c.Param("username")
	c.JSON(200, gin.H{
		"message": "Профиль пользователя " + username,
	})
}
