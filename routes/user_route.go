package routes

import (
	"go-mongodb/configs"
	"go-mongodb/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	userCollection := configs.GetCollection(configs.DB, "users")
	userHandler := controllers.NewUserHandler(userCollection)
	router.POST("/users", userHandler.CreateUser)
	router.PUT("/users/:userId", userHandler.EditAUser)
	router.GET("/users", userHandler.GetUsers)
	router.GET("/users/:userId", userHandler.GetAUser)
	router.DELETE("/users/:email", userHandler.DeleteUser)
}
