package job

import (
	"github.com/mrrizkin/crontastic/app/models"
)

type Service struct {
	repo *Repo
}

func NewService(repo *Repo) *Service {
	return &Service{repo}
}

func (s *Service) Create(job *models.Job) (*models.Job, error) {
	err := s.repo.Create(job)
	if err != nil {
		return nil, err
	}

	return job, nil
}

func (s *Service) FindAll() ([]models.Job, error) {
	jobs, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return jobs, nil
}

func (s *Service) FindByID(id uint) (*models.Job, error) {
	job, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return job, nil
}

func (s *Service) Update(id uint, job *models.Job) (*models.Job, error) {
	var err error

	_, err = s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = s.repo.Update(job)
	if err != nil {
		return nil, err
	}

	return job, nil
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}
