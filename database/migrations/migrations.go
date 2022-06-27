package migrations

import (
	"github.com/nogueirahy/api-gymfull/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Member{})
}
