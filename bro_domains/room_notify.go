package bro_domains

import (
	"time"

	"github.com/gofrs/uuid"
)

type RoomNotify struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;column:user_id" json:"user_id"`
	RoomName  string    `json:"room_name" gorm:"column:room_name" validate:"required"`
	Status    string    `json:"status" gorm:"column:status"`
	LineToken string    `json:"line_token" gorm:"column:line_token;index" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
