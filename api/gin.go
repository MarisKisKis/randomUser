package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"random_User/database"
	"random_User/model"
)

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get user",
	})
	return
}

func getContent() model.User {
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
	fmt.Printf("User Results: %v\n", data)
	return data
}

func convertResultsToUser(a database.User, b model.User) database.User {
	var userRes = b.Results[0]
	a.Gender = userRes.Gender
	a.Title = userRes.Name.Title
	a.FirstName = userRes.Name.First
	a.LastName = userRes.Name.Last
	a.StreetNumber = userRes.Location.Street.Number
	a.StreetName = userRes.Location.Street.Name
	a.City = userRes.Location.City
	a.State = userRes.Location.State
	a.Country = userRes.Location.Country
	a.Postcode = userRes.Location.Postcode
	a.Latitude = userRes.Location.Coordinates.Latitude
	a.Longitude = userRes.Location.Coordinates.Longitude
	a.TimezoneOffset = userRes.Location.Timezone.Offset
	a.TimezoneDescription = userRes.Location.Timezone.Description
	a.Email = userRes.Email
	a.Uuid = userRes.Login.Uuid
	a.Username = userRes.Login.Username
	a.Password = userRes.Login.Password
	a.Salt = userRes.Login.Salt
	a.Md5 = userRes.Login.Md5
	a.Sha1 = userRes.Login.Sha1
	a.Sha256 = userRes.Login.Sha256
	a.DobDate = userRes.Dob.Date
	a.DobAge = userRes.Dob.Age
	a.RegisteredDate = userRes.Registered.Date
	a.RegisteredAge = userRes.Registered.Age
	a.Phone = userRes.Phone
	a.Cell = userRes.Cell
	a.IdName = userRes.Id.Name
	a.IdValue = userRes.Id.Value
	a.PictureLarge = userRes.Picture.Large
	a.PictureMedium = userRes.Picture.Medium
	a.Thumbnail = userRes.Picture.Thumbnail
	a.Nat = userRes.Nat
	var userInf = b.Info
	a.Seed = userInf.Seed
	a.Results = userInf.Results
	a.Page = userInf.Page
	a.Version = userInf.Version

	return a
}

func postUser(c *gin.Context) {
	user1 := getContent()
	var user database.User
	user = convertResultsToUser(user, user1)
	/*var userName = userRes.Name[0]
	a.Title = userName.Title
	a.FirstName = userName.First
	a.LastName = userName.Last


	user.Title = userRes.Name.Title
	var userLoc = userRes.Location
	var street = userLoc.Street
	user.StreetNumber = street.Number
	user.StreetName = street.Name
	user.City = userLoc.City
	user.State = userLoc.State
	user.Country = userLoc.Country
	user.Postcode = userLoc.Postcode
	var userCord = userLoc.Coordinates
	user.Latitude = userCord.Latitude
	user.Longitude = userCord.Longitude
	var userTime = userLoc.Timezone
	user.TimezoneOffset = userTime.Offset
	user.TimezoneDescription = userTime.Description
	user.Email = userRes.Email
	user.Phone = userRes.Phone

	*/
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
