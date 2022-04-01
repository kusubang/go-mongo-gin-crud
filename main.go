package main

import (
	"go-mongodb/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.UserRoute(router)
	router.Run(":8080")
}
