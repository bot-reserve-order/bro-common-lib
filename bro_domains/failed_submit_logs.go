package bro_domains

import (
	"time"

	"github.com/gofrs/uuid"
)

type FailedSubmitLogs struct {
	ID           string    `json:"id" gorm:"primaryKey;column:id"`
	OwnerFleetId uuid.UUID `json:"ownerFleetId" gorm:"column:ownerFleetId"`
	CondId       int64     `json:"cond_id" gorm:"column:condId"`
	CarTypeText  string    `json:"car_type_text" gorm:"column:carTypeText"`
	Track        string    `json:"track" gorm:"column:track"`
	Code         int       `json:"code" gorm:"column:code"`
	Message      string    `json:"message" gorm:"column:message"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
