package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Krandr332/sait/controllers"
	"github.com/Krandr332/sait/routes"
)

func main() {
	r := gin.Default()

	// Подключение к базе данных или другие инициализации

	// Инициализация контроллеров
	homeController := &controllers.HomeController{}
	profileController := &controllers.ProfileController{}
	// Инициализация маршрутов
	routes.InitRoutes(r, homeController, profileController)

	r.Run(":8080")
}
