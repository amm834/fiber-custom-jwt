package database

import (
	"go-react-jwt/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// username:password@protocol(address)/dbname?param=value
	client, err := gorm.Open(mysql.Open("root:@/test"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// initialize the database
	DB = client

	client.AutoMigrate(&models.User{})
}
