// routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/your-username/social-network/controllers"
)

func InitRoutes(r *gin.Engine, homeController *controllers.HomeController, profileController *controllers.ProfileController) {
	r.GET("/", homeController.Home)
	r.GET("/profile/:username", profileController.GetProfile)
}
