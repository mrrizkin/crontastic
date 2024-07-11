package database

import (
	"github.com/mrrizkin/crontastic/app/models"
	"github.com/mrrizkin/crontastic/system/database"
)

type Repo struct {
	db *database.Database
}

func NewRepo(db *database.Database) *Repo {
	return &Repo{db}
}

func (r *Repo) Create(database *models.Database) error {
	return r.db.Create(database).Error
}

func (r *Repo) FindAll() ([]models.Database, error) {
	var databases []models.Database
	err := r.db.Find(&databases).Error
	return databases, err
}

func (r *Repo) FindByID(id uint) (*models.Database, error) {
	var database models.Database
	err := r.db.First(&database, id).Error
	return &database, err
}

func (r *Repo) Update(database *models.Database) error {
	return r.db.Save(database).Error
}

func (r *Repo) Delete(id uint) error {
	return r.db.Delete(&models.Database{}, id).Error
}
