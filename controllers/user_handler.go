package controllers

import (
	"go-mongodb/models"
	"go-mongodb/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	// collection *mongo.Collection
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func (h *UserHandler) GetAUser(c *gin.Context) {
	userId := c.Param("userId")

	user, err := h.service.GetAUser(userId)
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

	newUser, err := h.service.CreateAUser(&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "already exists"})
		return
	}

	c.JSON(http.StatusOK, newUser)
}

func (h UserHandler) EditAUser(c *gin.Context) {

	userId := c.Param("userId")
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	updatedUser, err := h.service.EditAUser(userId, &user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)

}

func (h *UserHandler) DeleteAUser(c *gin.Context) {

	userId := c.Param("userId")

	err := h.service.DeleteAUser(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User is deleted successfully"})
}
