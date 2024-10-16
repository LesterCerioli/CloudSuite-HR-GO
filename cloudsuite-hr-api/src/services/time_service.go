package services

import (
	"cloudsuite-hr-api/models"
	"cloudsuite-hr-api/repositories"
	"errors"
	"fmt"
	"log"
	"time"
	
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

// CreateTime with retry and logging
func (s *timeService) CreateTime(timeCreated models.Time) error {
	maxRetries := 3
	retryInterval := 3 * time.Millisecond
	var err error

	for i := 0; i < maxRetries; i++ {
		startDate := time.Now()
		log.Printf("Processing started at: %s", startDate)

		err = s.repo.CreateTime(timeCreated)
		endDate := time.Now()

		if err == nil {
			processingDuration := endDate.Sub(startDate)
			log.Printf("Success: Date: %s, Start: %s, End: %s, Duration: %s",
				time.Now().Format(time.RFC3339),
				startDate.Format(time.RFC3339),
				endDate.Format(time.RFC3339),
				processingDuration)
			return nil
		}

		processingDuration := endDate.Sub(startDate)
		log.Printf("Failure: Date: %s, Start: %s, End: %s, Duration: %s, Error: %s",
			time.Now().Format(time.RFC3339),
			startDate.Format(time.RFC3339),
			endDate.Format(time.RFC3339),
			processingDuration, err.Error())

		// Aguarda antes de tentar novamente
		time.Sleep(retryInterval)
	}
	return errors.New(fmt.Sprintf("Max retries reached: %v", err))
}

// GetAllTimes with retry and logging
func (s *timeService) GetAllTimes() ([]models.Time, error) {
	maxRetries := 3
	retryInterval := 3 * time.Millisecond
	var err error
	var times []models.Time

	for i := 0; i < maxRetries; i++ {
		startDate := time.Now()
		log.Printf("Processing started at: %s", startDate)

		times, err = s.repo.GetAllTimes()
		endDate := time.Now()

		if err == nil {
			processingDuration := endDate.Sub(startDate)
			log.Printf("Success: Date: %s, Start: %s, End: %s, Duration: %s",
				time.Now().Format(time.RFC3339),
				startDate.Format(time.RFC3339),
				endDate.Format(time.RFC3339),
				processingDuration)
			return times, nil
		}

		processingDuration := endDate.Sub(startDate)
		log.Printf("Failure: Date: %s, Start: %s, End: %s, Duration: %s, Error: %s",
			time.Now().Format(time.RFC3339),
			startDate.Format(time.RFC3339),
			endDate.Format(time.RFC3339),
			processingDuration, err.Error())

		// Aguarda antes de tentar novamente
		time.Sleep(retryInterval)
	}

	return nil, errors.New(fmt.Sprintf("Max retries reached: %v", err))
}

// GetTimesByDate with retry and logging
func (s *timeService) GetTimesByDate(date string) ([]models.Time, error) {
	maxRetries := 3
	retryInterval := 3 * time.Millisecond
	var err error
	var times []models.Time

	for i := 0; i < maxRetries; i++ {
		startDate := time.Now()
		log.Printf("Processing started at: %s", startDate)

		times, err = s.repo.GetTimesByDate(date)
		endDate := time.Now()

		if err == nil {
			processingDuration := endDate.Sub(startDate)
			log.Printf("Success: Date: %s, Start: %s, End: %s, Duration: %s",
				time.Now().Format(time.RFC3339),
				startDate.Format(time.RFC3339),
				endDate.Format(time.RFC3339),
				processingDuration)
			return times, nil
		}

		processingDuration := endDate.Sub(startDate)
		log.Printf("Failure: Date: %s, Start: %s, End: %s, Duration: %s, Error: %s",
			time.Now().Format(time.RFC3339),
			startDate.Format(time.RFC3339),
			endDate.Format(time.RFC3339),
			processingDuration, err.Error())

		// Aguarda antes de tentar novamente
		time.Sleep(retryInterval)
	}

	return nil, errors.New(fmt.Sprintf("Max retries reached: %v", err))
}
