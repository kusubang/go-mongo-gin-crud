package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	collection *mongo.Collection
}

func NewUserHandler(collection *mongo.Collection) *UserHandler {
	return &UserHandler{collection}
}

func (h *UserHandler) getUsersHandler(c *gin.Context) {
	// collection := client.Database("test").Collection("users")
	cursor, err := h.collection.Find(context.TODO(), bson.D{}, nil)

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

func (h *UserHandler) getUserHandler(c *gin.Context) {
	email := c.Param("email")
	var user User
	d := h.collection.FindOne(context.TODO(), gin.H{"email": email})

	d.Decode(&user)

	if err := d.Err(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "not exists"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) addUserHandler(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// var result primitive.M
	d := h.collection.FindOne(context.TODO(), bson.D{{"email", user.Email}})
	if err := d.Err(); err == nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "already exists"})
		return
	}

	_, err := h.collection.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) deleteUserHandler(c *gin.Context) {
	email := c.Param("email")

	res, err := h.collection.DeleteOne(context.TODO(), bson.D{{"email", email}})
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "fail to delete"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": res.DeletedCount})
}
