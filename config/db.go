package config

import (
	"fmt"
	"time"

	"example.com/sa-65-example/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("sa.db?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected database")
	db = database
}

func SetupDatabase() {

	// database, err := gorm.Open(sqlite.Open("sa.db?cache=shared"), &gorm.Config{})

	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// fmt.Println("connected database")

	// Migrate the schema

	db.AutoMigrate(
		&models.Users{},
		&models.Genders{},
	)

	GenderMale := models.Genders{Gender: "Male"}
	GenderFemale := models.Genders{Gender: "Female"}

	db.FirstOrCreate(&GenderMale, &models.Genders{Gender: "Male"})
	db.FirstOrCreate(&GenderFemale, &models.Genders{Gender: "Female"})

	hashedPassword, _ := HashPassword("123456")
	BirthDay, _ := time.Parse("2006-01-02", "1988-11-12")
	User := &models.Users{
		FirstName: "Software",
		LastName:  "Analysis",
		Email:     "sa@gmail.com",
		Age:       80,
		Password:  hashedPassword,
		BirthDay:  BirthDay,
		GenderID:  1,
	}
	db.FirstOrCreate(User, &models.Users{
		Email: "sa@gmail.com",
	})

}
