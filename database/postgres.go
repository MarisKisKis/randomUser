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
	"random_User/model"
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
	country             string `json:"country"`
	postcode            string `json:"postcode"`
	latitude            string `json:"latitude"`
	longitude           string `json:"longitude"`
	timezoneOffset      string `json:"offset"`
	timezoneDescription string `json:"description"`
	Email               string `json:"email"`
	uuid                string `json:"uuid"`
	username            string `json:"username"`
	password            string `json:"password"`
	salt                string `json:"salt"`
	md5                 string `json:"md5"`
	sha1                string `json:"sha1"`
	sha256              string `json:"sha256"`
	dobDate             string `json:"date"`
	dobAge              int    `json:"age"`
	registeredDate      string `json:"date"`
	registeredAge       int    `json:"age"`
	Phone               string `json:"phone"`
	cell                string `json:"cell"`
	idName              string `json:"name"`
	idValue             string `json:"value"`
	pictureLarge        string `json:"large"`
	pictureMedium       string `json:"medium"`
	thumbnail           string `json:"thumbnail"`
	nat                 string `json:"nat"`
	seed                string `json:"seed"`
	results             int    `json:"results"`
	page                int    `json:"page"`
	version             string `json:"version"`
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

func CreateUser(a model.User) (model.User, error) {
	res := db.Create(a)
	if res.RowsAffected == 0 {
		return model.User{}, errors.New("user not created")
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
