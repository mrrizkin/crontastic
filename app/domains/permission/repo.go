package permission

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

func (r *Repo) Create(permission *models.Permission) error {
	return r.db.Create(permission).Error
}

func (r *Repo) FindAll() ([]models.Permission, error) {
	var permissions []models.Permission
	err := r.db.Find(&permissions).Error
	return permissions, err
}

func (r *Repo) FindByID(id uint) (*models.Permission, error) {
	var permission models.Permission
	err := r.db.First(&permission, id).Error
	return &permission, err
}

func (r *Repo) Update(permission *models.Permission) error {
	return r.db.Save(permission).Error
}

func (r *Repo) Delete(id uint) error {
	return r.db.Delete(&models.Permission{}, id).Error
}
