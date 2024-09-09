package db

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// USERNAME = pkg.GetConfig("USERNAME")
	// PASSWORD = pkg.GetConfig("PASSWORD")
	// HOST     = pkg.GetConfig("HOST")
	// PORT     = pkg.GetConfig("PORT")
	// DBNAME   = pkg.GetConfig("DBNAME")
	USERNAME = "root"
	//PASSWORD = "indonesi4"
	//HOST     = "localhost"
	PASSWORD = "123456"
	HOST     = "141.11.25.60"
	PORT     = "3306"
	DBNAME   = "patungan_db"
)

func OpenDB() (*gorm.DB, error) {
	//return sql.Open("mysql", "root:indonesi4@tcp(localhost:3306)/brilian_db?parseTime=true")

	//naro plain aja biar ga ribet naro di env
	//username := "root"
	//password := "indonesi4"
	//host := "127.0.0.1"
	//port := "3306"
	//dbname := "brilian_db"

	logrus.Print(USERNAME, PASSWORD, HOST, PORT, DBNAME)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", USERNAME, PASSWORD, HOST, PORT, DBNAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatalln("error connect to database :", err.Error())
		return nil, err
	}

	return db, nil
}
