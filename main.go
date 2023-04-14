package main

import (
	"random_User/api"
	"random_User/database"

	_ "random_User/api"
	_ "random_User/database"
)

func init() {
	database.NewPostgreSQLClient()
}

func main() {
	r := api.SetupRouter()
	r.Run("localhost:8081")
}
