package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Dsn     string
	PortApi int64
	DB      *gorm.DB
)

func LoadConfig() {
	if erro := godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}
	port, erro := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 64)
	if erro != nil {
		log.Fatal(erro)
	}
	Dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), port, os.Getenv("DB_NAME"))
	PortApi, erro = strconv.ParseInt(os.Getenv("PORT"), 10, 64)
	if erro != nil {
		log.Fatal(erro)
	}
}

func ConnectDB() {
	var erro error
	DB, erro = gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if erro != nil {
		log.Fatal(erro)
	}
}
