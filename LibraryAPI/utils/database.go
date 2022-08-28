package utils

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func migrateDB(db *gorm.DB) (err error) {
	err = db.AutoMigrate(&Book{})
	return err
}

func ConnectDatabase() (*gorm.DB, error) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Istanbul", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DATABASE"), os.Getenv("POSTGRES_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err

}

func InitDatabase() {
	db, err := ConnectDatabase()
	if err != nil {
		fmt.Println("Database connection error!")
	} else {
		fmt.Println("Database connection success!")
	}

	err = migrateDB(db)

	if err != nil {
		fmt.Println("Database Migrate error!")
	} else {
		fmt.Println("Database migrated!")
	}

}
