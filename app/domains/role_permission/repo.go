package role_permission

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

func (r *Repo) Create(role_permission *models.RolePermission) error {
	return r.db.Create(role_permission).Error
}

func (r *Repo) FindAll() ([]models.RolePermission, error) {
	var role_permissions []models.RolePermission
	err := r.db.Find(&role_permissions).Error
	return role_permissions, err
}

func (r *Repo) FindByID(id uint) (*models.RolePermission, error) {
	var role_permission models.RolePermission
	err := r.db.First(&role_permission, id).Error
	return &role_permission, err
}

func (r *Repo) Update(role_permission *models.RolePermission) error {
	return r.db.Save(role_permission).Error
}

func (r *Repo) Delete(id uint) error {
	return r.db.Delete(&models.RolePermission{}, id).Error
}
