package models

import "time"

type Prediction struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	UserID          uint      `gorm:"not null;uniqueIndex:idx_competitor_day" json:"user_id"`
	TournamentID    uint      `gorm:"not null;index" json:"tournament_id"`
	TournamentDayID uint      `gorm:"not null;uniqueIndex:idx_competitor_day" json:"tournament_day_id"`
	SubmittedAt     time.Time `json:"submitted_at"`
}
