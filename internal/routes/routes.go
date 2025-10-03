package routes

import (
	"farm-investment/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine,
	userHandler *handler.UserHandler,
) {
	// User
	r.GET("/users", userHandler.GetUsers)
	r.POST("/users", userHandler.CreateUser)
}
