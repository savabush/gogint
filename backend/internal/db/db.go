package db

import (
	"fmt"
	"github.com/savabush/blog/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func connect() *gorm.DB {
	databaseConfig := config.Settings.Database
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Krasnoyarsk",
		databaseConfig.Host,
		databaseConfig.User,
		databaseConfig.Password,
		databaseConfig.Name,
		databaseConfig.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

var DB = connect()
