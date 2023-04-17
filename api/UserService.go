package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"random_User/database"
	"random_User/model"
)

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
