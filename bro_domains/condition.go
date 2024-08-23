package bro_domains

import (
	"time"

	"github.com/bot-reserve-order/bro-common-lib/bro_enum"
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

type Conditions struct {
	ID              uint             `gorm:"primaryKey" json:"id"`
	Status          string           `json:"status"`
	Area            string           `json:"area"`
	Cartype         string           `json:"car_type"`
	Track           string           `json:"track"`
	RoomNotifyID    uint             `json:"room_notify_id"`
	MaxRange        uint             `json:"max_range"`
	SerialID        string           `json:"serial_id"`
	OrderDetail     Order            `gorm:"type:jsonb;" json:"order_detail"`
	UserID          uuid.UUID        `gorm:"type:uuid;" json:"user_id"`
	UserRootID      uuid.UUID        `gorm:"type:uuid;" json:"user_root_id"`
	FounderUserID   uuid.UUID        `gorm:"type:uuid;" json:"founder_user_id"`
	ReferenceUserID uuid.UUID        `gorm:"type:uuid;" json:"reference_user_id"`
	Credits         []uint           `json:"credits"`
	Boost           decimal.Decimal  `json:"boost"`
	PackageID       bro_enum.Package `json:"package_id"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
