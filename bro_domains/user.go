package bro_domains

import (
	"time"

	"github.com/bot-reserve-order/bro-common-lib/bro_enum"
	"github.com/gofrs/uuid"
)

type User struct {
	ID          uuid.UUID         `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	CustomID    string            `gorm:"uniqueIndex" json:"custom_id"`
	DisplayName string            `json:"display_name" gorm:"column:display_name" validate:"required"`
	Role        bro_enum.RoleUser `json:"role" gorm:"column:role"`
	Username    string            `json:"username" gorm:"column:username" validate:"required"`
	Password    string            `json:"password" gorm:"column:password" validate:"required"`
	Status      bro_enum.Status   `json:"status" gorm:"column:status"`
	Parent      uuid.UUID         `json:"parent" gorm:"column:parent"`
	CompanyID   uint              `json:"company_id" gorm:"column:company_id"`
	StaffID     uint              `json:"staff_id" gorm:"column:staff_id"`
	Reference   string            `json:"reference" gorm:"column:reference"`
	PackageID   bro_enum.Package  `json:"package_id" gorm:"column:package_id"`
	FleetToken  string            `json:"fleet_token" gorm:"column:fleet_token"`
	LineToken   string            `json:"line_token" gorm:"column:line_token"`
	TrueWallet  string            `json:"true_wallet" gorm:"column:true_wallet"`
	RunBot      bool              `json:"runbot" gorm:"column:runbot;default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
