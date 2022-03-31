package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MONGO_URI = "mongodb://1234:1234@127.0.0.1:27017/"

func initMongo(uri string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	userCollection := client.Database("test").Collection("users")
	userHandler := NewUserHandler(userCollection)
	router.GET("/users", userHandler.getUsersHandler)
	router.GET("/users/:email", userHandler.getUserHandler)
	router.POST("/users", userHandler.addUserHandler)
	router.DELETE("/users/:email", userHandler.deleteUserHandler)

	router.Run("localhost:8080")
	return router
}

var client *mongo.Client

func main() {
	client = initMongo(MONGO_URI)
	defer client.Disconnect(context.Background())

	r := setupRouter()
	r.Run()

}
