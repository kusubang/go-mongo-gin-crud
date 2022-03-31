package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func getUsersHandler(c *gin.Context) {
	collection := client.Database("test").Collection("users")
	cursor, err := collection.Find(context.TODO(), bson.D{}, nil)

	var users []*User
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	for cursor.Next(context.TODO()) {
		var elem User
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, &elem)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	cursor.Close(context.TODO())

	c.IndentedJSON(http.StatusOK, users)
}

func getUserHandler(c *gin.Context) {
	email := c.Param("email")
	collection := client.Database("test").Collection("users")
	var user User
	d := collection.FindOne(context.TODO(), gin.H{"email": email})

	d.Decode(&user)

	if err := d.Err(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "not exists"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func addUserHandler(c *gin.Context) {
	var user User

	collection := client.Database("test").Collection("users")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// var result primitive.M
	d := collection.FindOne(context.TODO(), bson.D{{"email", user.Email}})
	if err := d.Err(); err == nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "already exists"})
		return
	}

	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func deleteUserHandler(c *gin.Context) {
	email := c.Param("email")
	collection := client.Database("test").Collection("users")

	res, err := collection.DeleteOne(context.TODO(), bson.D{{"email", email}})
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "fail to delete"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": res.DeletedCount})

	// json.NewEncoder(w).Encode(res.DeletedCount) // return number of

}
