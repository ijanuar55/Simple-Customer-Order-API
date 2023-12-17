package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectMYSQL() *gorm.DB {
	Database = SetupDB()
	return Database
}

func SetupDB() *gorm.DB {
	var err error
	host := os.Getenv("DB_HOST_MYSQL")
	username := os.Getenv("DB_USER_MYSQL")
	password := os.Getenv("DB_PASSWORD_MYSQL")
	databaseName := os.Getenv("DB_NAME_MYSQL")
	port := os.Getenv("DB_PORT_MYSQL")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, databaseName)
	Database1, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Successfully connected to the database")
	}
	return Database1
}
