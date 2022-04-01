package routes

import (
	"go-mongodb/configs"
	"go-mongodb/controllers"
	"go-mongodb/services"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	userCollection := configs.GetCollection(configs.DB, "users")
	userService := services.NewUserServiceMongo(userCollection)
	// userService := services.NewUserServiceDummy()
	userHandler := controllers.NewUserHandler(userService)
	router.POST("/users", userHandler.CreateUser)
	router.PUT("/users/:userId", userHandler.EditAUser)
	router.GET("/users", userHandler.GetAllUsers)
	router.GET("/users/:userId", userHandler.GetAUser)
	router.DELETE("/users/:userId", userHandler.DeleteAUser)
}
