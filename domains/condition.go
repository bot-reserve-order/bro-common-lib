package domains

import (
	"time"

	"github.com/gofrs/uuid"
)

type Conditions struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Status          string    `json:"status"`
	Area            string    `json:"area"`
	Cartype         string    `json:"car_type"`
	Track           string    `json:"track"`
	RoomNotifyID    uint      `json:"room_notify_id"`
	MaxRange        uint      `json:"max_range"`
	SerialID        string    `json:"serial_id"`
	OrderDetail     BaseInfo  `gorm:"type:jsonb;" json:"order_detail"`
	UserID          uuid.UUID `gorm:"type:uuid;" json:"user_id"`
	UserRootID      uuid.UUID `gorm:"type:uuid;" json:"user_root_id"`
	FounderUserID   uuid.UUID `gorm:"type:uuid;" json:"founder_user_id"`
	ReferenceUserID uuid.UUID `gorm:"type:uuid;" json:"reference_user_id"`
	CreditID        uint      `json:"credit_id"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
