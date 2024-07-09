package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("Hello World!")
// }

// func main() {
// 	r := gin.New()

// 	r.GET("/users", listUsersHandler)
// 	r.POST("/users", createUserHandler)
// 	r.DELETE("/users/:id", deleteUserHandler)

// 	r.Run()
// }

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var users = []User{
	{ID: "1", Name: "test no.1", Description: "King"},
	{ID: "2", Name: "test no.2", Description: "Queen"},
	{ID: "3", Name: "test no.3", Description: "Jack"},
}

// func listUsersHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, books)
// }

// func createUserHandler(c *gin.Context) {
// 	var user User

// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	users = append(users, user)

// 	c.JSON(http.StatusCreated, user)
// }

// func deleteUserHandler(c *gin.Context) {
// 	id := c.Param("id")

// 	for i, a := range users {
// 		if a.ID == id {
// 			users = append(users[:i], users[i+1:]...)
// 			break
// 		}
// 	}

// 	c.Status(http.StatusNoContent)
// }

func main() {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	r.POST("/users", func(c *gin.Context) {
		var user User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		users = append(users, user)
		c.JSON(http.StatusCreated, user)
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")

		for i, a := range users {
			if a.ID == id {
				users = append(users[:i], users[i+1:]...)
				break
			}
		}
		c.Status(http.StatusNoContent)
	})

	r.Run()
}
