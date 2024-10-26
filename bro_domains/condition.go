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
	ID              uint                      `gorm:"primaryKey" json:"id" firestore:"id"`
	Status          string                    `json:"status" gorm:"column:status" firestore:"status"`
	Area            string                    `json:"area" gorm:"column:area" validate:"required" firestore:"area"`
	CarTypeText     string                    `json:"car_type_text" gorm:"column:car_type_text" validate:"required" firestore:"car_type_text"`
	Track           pq.StringArray            `json:"track" gorm:"column:track;type:text[]" validate:"required" firestore:"track"`
	GroupNotify     uuid.UUID                 `json:"group_notify_id" gorm:"column:group_notify_id" firestore:"group_notify_id"`
	MaxRange        uint                      `json:"max_range" gorm:"column:max_range" firestore:"max_range"`
	SerialID        string                    `json:"serial_id" gorm:"column:serial_id;index" firestore:"serial_id"`
	OrderDetail     datatypes.JSONType[Order] `json:"order_detail" gorm:"type:jsonb;column:order_detail" firestore:"order_detail"`
	UserID          uuid.UUID                 `json:"user_id" gorm:"type:uuid;column:user_id;index" firestore:"user_id"`
	UserRootID      uuid.UUID                 `json:"user_root_id" gorm:"type:uuid;column:user_root_id;index" firestore:"user_root_id"`
	FounderUserID   uuid.UUID                 `json:"founder_user_id" gorm:"type:uuid;column:founder_user_id;index" firestore:"founder_user_id"`
	ReferenceUserID uuid.UUID                 `json:"reference_user_id" gorm:"type:uuid;column:reference_user_id;index" firestore:"reference_user_id"`
	Credits         pq.Int64Array             `json:"credits" gorm:"column:credits;type:integer[]" firestore:"credits"`
	Boost           decimal.Decimal           `json:"boost" firestore:"boost"`
	PackageID       bro_enum.Package          `json:"package_id" firestore:"package_id"`
	Note            string                    `json:"note" firestore:"note"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
