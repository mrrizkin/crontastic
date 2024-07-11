package job

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

func (r *Repo) Create(job *models.Job) error {
	return r.db.Create(job).Error
}

func (r *Repo) FindAll() ([]models.Job, error) {
	var jobs []models.Job
	err := r.db.Find(&jobs).Error
	return jobs, err
}

func (r *Repo) FindByID(id uint) (*models.Job, error) {
	var job models.Job
	err := r.db.First(&job, id).Error
	return &job, err
}

func (r *Repo) Update(job *models.Job) error {
	return r.db.Save(job).Error
}

func (r *Repo) Delete(id uint) error {
	return r.db.Delete(&models.Job{}, id).Error
}
