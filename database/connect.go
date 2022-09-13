package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/BangsaMU/config"
	"github.com/BangsaMU/internals/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Idiot")
	}

	// Connection URL to connect to mysql Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port, config.Config("DB_NAME"))
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details

	log.Println(dsn)
	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(mysql.Open(dsn))
	// DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Info),
	// })

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	// Migrate the database
	DB.AutoMigrate(&model.Note{})
	fmt.Println("Database Migrated")
}
