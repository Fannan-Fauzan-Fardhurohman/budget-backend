package main

import (
	"bougette-backend/common"
	"bougette-backend/internal/models"
	"log"
)

func main() {
	db, err := common.NewMysql()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.UserModel{})
	if err != nil {
		panic(err)
	}
	log.Println("Migration completed")
}
