package seed

import (
	"log"

	"github.com/AnibrataMalkhandi/Google-Map-Leads-Generation/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {

	err := db.AutoMigrate(
		&models.UserDetails{},
	)

	if err != nil {
		log.Fatalln(err)
		return
	}
}