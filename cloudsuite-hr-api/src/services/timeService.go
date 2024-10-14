package services

import (
	"cloudsuite-hr-api/models"
	"cloudsuite-hr-api/repositories"
)

type TimeService interface {
	CreateTime(time models.Time) error
	GetAllTimes() ([]models.Time, error)
	GetTimesByDate(date string) ([]models.Time, error)
}

type timeService struct {
	repo repositories.TimeRepository
}

func NewTimeService(repo repositories.TimeRepository) TimeService {
	return &timeService{repo: repo}
}

func (s *timeService) CreateTime(time models.Time) error {
	return s.repo.CreateTime(time)
}

func (s *timeService) GetAllTimes() ([]models.Time, error) {
	return s.repo.GetAllTimes()
}

func (s *timeService) GetTimesByDate(date string) ([]models.Time, error) {
	return s.repo.GetTimesByDate(date)
}
