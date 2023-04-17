package database

import (
	"errors"
	_ "errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *gorm.DB
var err error

type User struct {
	ID                  uint   `json:"id" gorm:"primary_key"`
	Gender              string `json:"gender"`
	Title               string `json:"title"`
	FirstName           string `json:"first"`
	LastName            string `json:"last"`
	StreetNumber        int    `json:"number"`
	StreetName          string `json:"name"`
	City                string `json:"city"`
	State               string `json:"state"`
	Country             string `json:"country"`
	Postcode            int32  `json:"postcode"`
	Latitude            string `json:"latitude"`
	Longitude           string `json:"longitude"`
	TimezoneOffset      string `json:"offset"`
	TimezoneDescription string `json:"description"`
	Email               string `json:"email"`
	Uuid                string `json:"uuid"`
	Username            string `json:"username"`
	Password            string `json:"password"`
	Salt                string `json:"salt"`
	Md5                 string `json:"md5"`
	Sha1                string `json:"sha1"`
	Sha256              string `json:"sha256"`
	DobDate             string `json:"date"`
	DobAge              int    `json:"age"`
	RegisteredDate      string `json:"date"`
	RegisteredAge       int    `json:"age"`
	Phone               string `json:"phone"`
	Cell                string `json:"cell"`
	IdName              string `json:"name"`
	IdValue             string `json:"value"`
	PictureLarge        string `json:"large"`
	PictureMedium       string `json:"medium"`
	Thumbnail           string `json:"thumbnail"`
	Nat                 string `json:"nat"`
	Seed                string `json:"seed"`
	Results             int    `json:"results"`
	Page                int    `json:"page"`
	Version             string `json:"version"`
}

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	return os.Getenv(key)
}

func NewPostgreSQLClient() {
	var (
		host     = getEnvVariable("DB_HOST")
		port     = getEnvVariable("DB_PORT")
		user     = getEnvVariable("DB_USER")
		dbname   = getEnvVariable("DB_NAME")
		password = getEnvVariable("DB_PASSWORD")
	)

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
	)

	db, err = gorm.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(User{})

}

func CreateUser(a *User) (*User, error) {
	res := db.Create(a)
	if res.RowsAffected == 0 {
		return &User{}, errors.New("user not created")
	}
	return a, nil
}

func GetUser(id string) (*User, error) {
	var user User
	res := db.First(&user, id)
	if res.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func GetAllUsers() ([]*User, error) {
	var users []*User
	res := db.Find(&users)
	if res.Error != nil {
		return nil, errors.New("authors not found")
	}
	return users, nil
}

func DeleteUser(id string) error {
	var deleteUser User
	result := db.Where(id).Delete(&deleteUser)
	if result.RowsAffected == 0 {
		return errors.New("user data not deleted")
	}
	return nil
}
