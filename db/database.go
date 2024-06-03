package db

import (
	"fmt"
	"os"

	"gitlab.com/Std217/test/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToMariaDB() (*gorm.DB, error) {
	if err := config.EnvLoad(); err != nil {
		println(err)
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbURL := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")
	Path := dbUser + ":" + dbPass + "@tcp(" + dbURL + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
	println(Path)
	db, err := gorm.Open(mysql.Open(Path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MariaDB!")
	return db, nil
}
