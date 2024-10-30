package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	router.POST("/users", func(c *gin.Context){
		var user User

		err := c.ShouldBindJSON(&user)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User created successfully",
			"user": user,
		})
	})

	router.Run(":8080")
}