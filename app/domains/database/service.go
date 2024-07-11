package database

import (
	"github.com/mrrizkin/crontastic/app/models"
)

type Service struct {
	repo *Repo
}

func NewService(repo *Repo) *Service {
	return &Service{repo}
}

func (s *Service) Create(database *models.Database) (*models.Database, error) {
	err := s.repo.Create(database)
	if err != nil {
		return nil, err
	}

	return database, nil
}

func (s *Service) FindAll() ([]models.Database, error) {
	databases, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return databases, nil
}

func (s *Service) FindByID(id uint) (*models.Database, error) {
	database, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return database, nil
}

func (s *Service) Update(id uint, database *models.Database) (*models.Database, error) {
	var err error

	_, err = s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = s.repo.Update(database)
	if err != nil {
		return nil, err
	}

	return database, nil
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}
