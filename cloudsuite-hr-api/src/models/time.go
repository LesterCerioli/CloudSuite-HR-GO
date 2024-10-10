package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Time struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Date           string    `json:"Date"`
	EntryTime      time.Time `json:"entry_time"`
	LunchEntryTime time.Time `json:"lunch_entry_time"`
	LunchExitTime  time.Time `json:"lunch_exit_time"`
	ExitTime       time.Time `json:"exit_time"`
}

func (t *Time) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return
}
