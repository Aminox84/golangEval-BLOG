package routes

import (
	"blog/internal/models"
	"blog/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userService *services.UserService
}

func NewUserRoutes(userService *services.UserService) *UserRoutes {
	return &UserRoutes{userService}
}

func (ur *UserRoutes) RegisterRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/users", ur.createUser)

	}
}

func (ur *UserRoutes) createUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := ur.userService.CreateUser(&user)
	if err != nil {
		log.Printf("Error creating user: %v", err) // Log the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user!"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"userID": userID, "Lastname": user.Lastname, "Email": user.Email})
}
