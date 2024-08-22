package routes

import (
	"wachirawut_agnos_backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func GroupRoutes(router *gin.Engine, controller *controllers.Controller) {
	routes := router.Group("/api")
	{
		routes.POST("/strong_password_steps", controller.StrongPasswordStep)
	}
}
