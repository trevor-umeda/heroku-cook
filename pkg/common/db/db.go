package db

import (
	"github.com/trevor-umeda/heroku-cook/pkg/common/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Recipe{})
	db.AutoMigrate(&models.Tag{})
	return db
}
