package main

import (
	"singleservice/initializers"
	model "singleservice/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&model.Conversation{}, &model.Chat{}, &model.User{})
}
