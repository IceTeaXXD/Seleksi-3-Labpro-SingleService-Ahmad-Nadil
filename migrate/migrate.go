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
	initializers.DB.AutoMigrate(&model.User{}, &model.Perusahaan{}, &model.Barang{})
}
