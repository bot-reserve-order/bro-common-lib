package domains

import (
	"time"

	"github.com/gofrs/uuid"
)

type RoomNotify struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;" json:"user_id"`
	RoomName  string    `json:"room_name"`
	Status    string    `json:"status"`
	LineToken string    `json:"line_token"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
