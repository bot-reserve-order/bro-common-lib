package bro_domains

import (
	"time"

	"github.com/bot-reserve-order/bro-common-lib/bro_enum"
	"github.com/gofrs/uuid"
)

type User struct {
	ID          uuid.UUID         `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	CustomID    string            `gorm:"uniqueIndex" json:"custom_id"`
	DisplayName string            `json:"display_name"`
	Role        bro_enum.RoleUser `json:"role"`
	Username    string            `json:"username"`
	Password    string            `json:"password"`
	Status      bro_enum.Status   `json:"status"`
	Parent      uuid.UUID         `json:"parent"`
	CompanyID   uint              `json:"company_id"`
	StaffID     uint              `json:"staff_id"`
	Reference   string            `json:"reference"`
	PackageID   bro_enum.Package  `json:"package_id"`
	FleetToken  string            `json:"fleet_token"`
	LineToken   string            `json:"line_token"`
	TrueWallet  string            `json:"true_wallet"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
