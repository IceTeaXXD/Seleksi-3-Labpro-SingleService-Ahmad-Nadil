package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")
	DB , err = gorm.Open(postgres.Open("host=" + dbHost + " user=" + dbUsername + " password=" + dbPassword + " dbname=" + dbTable + " port=" + dbPort + " sslmode=disable TimeZone=Asia/Jakarta"), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	log.Println("database connection successful")
}
