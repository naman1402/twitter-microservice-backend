package database

import (
	"log"

	"github.com/naman1402/twitter-microservice-backend/auth/infra/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartPostgres() *gorm.DB {

	db_url := "postgres://postgres:postgres@localhost:5432/twitter_clone"
	db, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	// add user model
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.UserFollowers{})

	return db
}
