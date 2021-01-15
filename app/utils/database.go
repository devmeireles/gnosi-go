package utils

import (
	"fmt"

	"github.com/devmeireles/gnosi-api/app/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	err error
)

// Database struct
type Database struct {
	*gorm.DB
}

// InitDatabase inits de database
func InitDatabase(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var db = DB

	// DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	// db, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})

	db, err = gorm.Open(sqlite.Open("./database/gorm.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println("db err: ", err)
	}

	if err != nil {
		fmt.Println("db err: ", err)
	}

	DB = db

	migration()

	// db.Preload("Episodes").Find()
	// db.Preload("Seasons").Preload("Episodes").Find(&models.Catalogue{})
	// db.Preload("Seasons").Find(&models.Catalogue{})

	// db.Preload("Seasons").Find(&[]models.Catalogue{})

}

func migration() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Category{})
	DB.AutoMigrate(&models.Catalogue{})
	DB.AutoMigrate(&models.Season{})
	DB.AutoMigrate(&models.Episode{})
}

func DBConn() *gorm.DB {
	return DB
}
