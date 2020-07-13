package migrate

import (
	"github.com/DFrancis84/fantasyFootballDraftGo/internal/db/migrate/models"
	"github.com/jinzhu/gorm"
)

type Migrate struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *Migrate {
	return &Migrate{
		DB: db,
	}
}

func (m *Migrate) Migrate() {
	m.DB.AutoMigrate(&models.Pick{})
}
