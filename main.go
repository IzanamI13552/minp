package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// /time
	r.GET("/time", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"time": time.Now().Format("2006-01-02 15:04:05"),
		})
	})

	users := []User{
		{ID: 1, Name: "Ivan"},
		{ID: 2, Name: "Sasha"},
		{ID: 3, Name: "Kolya"},
	}

	// /users
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	r.Run(":8080")
}
