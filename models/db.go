package models

import (
	"fmt"
	"gihub.com/moeen/salamantex_back/config"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

func InitDB() {
	var err error

	pgUri := fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable",
		config.GetConfig().Postgres.User, config.GetConfig().Postgres.Pass,
		config.GetConfig().Postgres.Host, config.GetConfig().Postgres.DB)
	db, err = gorm.Open("postgres", pgUri)
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&User{}, &Transaction{})
}

func GetDB() *gorm.DB {
	return db
}
