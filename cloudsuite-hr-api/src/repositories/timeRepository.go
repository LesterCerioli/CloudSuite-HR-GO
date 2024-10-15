package repositories

import (
	"cloudsuite-hr-api/models"
	"gorm.io/gorm"
)

type TimeRepository interface {
	CreateTime(time models.Time) error
	GetAllTimes() ([]models.Time, error)
	GetTimesByDate(date string) ([]models.Time, error)
}

type timeRepository struct {
	db *gorm.DB
}

func NewTimeRepository(db *gorm.DB) TimeRepository {
	return &timeRepository{db: db}
}

func (r *timeRepository) CreateTime(time models.Time) error {
	return r.db.Create(&time).Error
}

func (r *timeRepository) GetAllTimes() ([]models.Time, error) {
	var times []models.Time
	err := r.db.Find(&times).Error
	return times, err
}

func (r *timeRepository) GetTimesByDate(date string) ([]models.Time, error) {
	var times []models.Time
	err := r.db.Where("date = ?", date).Find(&times).Error
	return times, err
}
