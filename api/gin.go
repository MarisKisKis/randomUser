package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"random_User/database"
)

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get user",
	})
	return
}

func postUser(c *gin.Context) {
	user1 := getContent()
	var user database.User
	user = convertResultsToUser(user, user1)
	res2, err2 := database.CreateUser(&user)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err2,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": res2,
	})
	return
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	user, err := database.GetUser(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
	return
}

func getUsers(c *gin.Context) {
	users, err := database.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	err := database.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
	})
	return
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", home)
	r.GET("/api/v1/users/:id", getUser)
	r.GET("/api/v1/users", getUsers)
	r.POST("/api/v1/users", postUser)
	r.DELETE("/api/v1/users/:id", deleteUser)
	return r
}
