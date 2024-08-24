package bro_domains

import (
	"time"

	"github.com/bot-reserve-order/bro-common-lib/bro_enum"
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

type Conditions struct {
	ID              uint             `gorm:"primaryKey" json:"id"`
	Status          string           `json:"status" grom:"column:status"`
	Area            string           `json:"area" grom:"column:area" validate:"required"`
	Cartype         string           `json:"car_type" grom:"column:car_type" validate:"required"`
	Track           string           `json:"track" grom:"column:track;index" validate:"required"`
	RoomNotifyID    uint             `json:"room_notify_id" grom:"column:room_notify_id"`
	MaxRange        uint             `json:"max_range" grom:"column:max_range"`
	SerialID        string           `json:"serial_id" grom:"column:serial_id;index"`
	OrderDetail     Order            `json:"order_detail" grom:"type:jsonb;column:order_detail"`
	UserID          uuid.UUID        `json:"user_id" grom:"type:uuid;column:user_id;index"`
	UserRootID      uuid.UUID        `json:"user_root_id" grom:"type:uuid;column:user_root_id;index"`
	FounderUserID   uuid.UUID        `json:"founder_user_id" grom:"type:uuid;column:founder_user_id;index"`
	ReferenceUserID uuid.UUID        `json:"reference_user_id" grom:"type:uuid;column:reference_user_id;index"`
	Credits         []uint           `json:"credits" grom:"column:credits"`
	Boost           decimal.Decimal  `json:"boost"`
	PackageID       bro_enum.Package `json:"package_id"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
