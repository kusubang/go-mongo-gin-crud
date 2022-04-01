package controllers

import (
	"context"
	"go-mongodb/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	collection *mongo.Collection
}

func NewUserHandler(collection *mongo.Collection) *UserHandler {
	return &UserHandler{collection}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	// collection := client.Database("test").Collection("users")
	cursor, err := h.collection.Find(context.TODO(), bson.D{}, nil)

	var users []models.User
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	for cursor.Next(context.TODO()) {
		var elem models.User
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, elem)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	cursor.Close(context.TODO())

	c.IndentedJSON(http.StatusOK, gin.H{"result": users})
}

func (h *UserHandler) GetAUser(c *gin.Context) {
	userId := c.Param("userId")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)
	err := h.collection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "not exists"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User

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

	newUser := models.User{
		Id:       primitive.NewObjectID(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	_, err := h.collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h UserHandler) EditAUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("userId")
	var user models.User
	defer cancel()
	objId, _ := primitive.ObjectIDFromHex(userId)

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	update := bson.M{"email": user.Email, "name": user.Name}
	result, err := h.collection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	//get updated user details
	var updatedUser models.User
	if result.MatchedCount == 1 {
		err := h.collection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, updatedUser)

}

func (h *UserHandler) DeleteAUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("userId")
	log.Printf("Delete userId: %s", userId)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	result, err := h.collection.DeleteOne(ctx, bson.M{"id": objId})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if result.DeletedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "User with specified ID not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User is deleted successfully"})
}
