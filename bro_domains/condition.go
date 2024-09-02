package bro_domains

import (
	"time"

	"github.com/bot-reserve-order/bro-common-lib/bro_enum"
	"github.com/gofrs/uuid"
	"github.com/lib/pq"
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
)

type Conditions struct {
	ID              uint                      `gorm:"primaryKey" json:"id"`
	Status          string                    `json:"status" gorm:"column:status"`
	Area            string                    `json:"area" gorm:"column:area" validate:"required"`
	CarTypeText     string                    `json:"car_type_text" gorm:"column:car_type_text" validate:"required"`
	Track           pq.StringArray            `json:"track" gorm:"column:track;type:text[]" validate:"required"`
	RoomNotifyID    uint                      `json:"room_notify_id" gorm:"column:room_notify_id"`
	MaxRange        uint                      `json:"max_range" gorm:"column:max_range"`
	SerialID        string                    `json:"serial_id" gorm:"column:serial_id;index"`
	OrderDetail     datatypes.JSONType[Order] `json:"order_detail" gorm:"type:jsonb;column:order_detail"`
	UserID          uuid.UUID                 `json:"user_id" gorm:"type:uuid;column:user_id;index"`
	UserRootID      uuid.UUID                 `json:"user_root_id" gorm:"type:uuid;column:user_root_id;index"`
	FounderUserID   uuid.UUID                 `json:"founder_user_id" gorm:"type:uuid;column:founder_user_id;index"`
	ReferenceUserID uuid.UUID                 `json:"reference_user_id" gorm:"type:uuid;column:reference_user_id;index"`
	Credits         pq.Int64Array             `json:"credits" gorm:"column:credits;type:integer[]"`
	Boost           decimal.Decimal           `json:"boost"`
	PackageID       bro_enum.Package          `json:"package_id"`
	Note            string                    `json:"note"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
