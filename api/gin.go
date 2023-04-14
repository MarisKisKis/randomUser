package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"random_User/database"
	"random_User/model"
	"time"
)

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get user",
	})
	return
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func get_content() model.User {
	// json data
	url := "https://randomuser.me/api/"
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	var data model.User
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Results: %v\n", data)
	return data
}
func postUser(c *gin.Context) {
	user1 := get_content()
	//	getJson("https://randomuser.me/api/", user1)
	//var user database.User
	//user.Gender =
	res2, err2 := database.CreateUser(user1)
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

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", home)
	r.GET("/api/v1/users/:id", getUser)
	r.POST("/api/v1/users", postUser)
	return r
}