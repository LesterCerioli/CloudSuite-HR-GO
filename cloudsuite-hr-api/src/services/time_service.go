package services

import (
	"cloudsuite-hr-api/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

type TimeService interface {
	CreateTime(time models.Time, db *gorm.DB) error
	GetAllTimes(db *gorm.DB) ([]models.Time, error)
	GetTimesByDate(date string, db *gorm.DB) ([]models.Time, error)
}

type timeService struct{}

func NewTimeService() TimeService {
	return &timeService{}
}

// CreateTime with retry and logging
func (s *timeService) CreateTime(timeCreated models.Time, db *gorm.DB) error {
	maxRetries := 3
	retryInterval := 3 * time.Millisecond
	var err error

	for i := 0; i < maxRetries; i++ {
		startDate := time.Now()
		log.Printf("Processing started at: %s", startDate)

		err = db.Create(&timeCreated).Error
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
func (s *timeService) GetAllTimes(db *gorm.DB) ([]models.Time, error) {
	maxRetries := 3
	retryInterval := 3 * time.Millisecond
	var err error
	var times []models.Time

	for i := 0; i < maxRetries; i++ {
		startDate := time.Now()
		log.Printf("Processing started at: %s", startDate)

		err = db.Find(&times).Error
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
func (s *timeService) GetTimesByDate(date string, db *gorm.DB) ([]models.Time, error) {
	maxRetries := 3
	retryInterval := 3 * time.Millisecond
	var err error
	var times []models.Time

	for i := 0; i < maxRetries; i++ {
		startDate := time.Now()
		log.Printf("Processing started at: %s", startDate)

		err = db.Where("Date = ?", date).Find(&times).Error
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
