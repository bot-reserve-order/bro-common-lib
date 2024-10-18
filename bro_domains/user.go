package bro_domains

import (
	"time"

	"github.com/bot-reserve-order/bro-common-lib/bro_enum"
	"github.com/gofrs/uuid"
)

type User struct {
	ID              uuid.UUID        `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Status          bro_enum.Status  `json:"status" gorm:"column:status"`
	Parent          uuid.UUID        `json:"parent" gorm:"column:parent"`
	Reference       uuid.UUID        `json:"reference" gorm:"column:reference"`
	PackageID       bro_enum.Package `json:"package_id" gorm:"column:package_id"`
	TrueWallet      string           `json:"true_wallet" gorm:"column:true_wallet"`
	RunBot          bool             `json:"runbot" gorm:"column:runbot;default:true"`
	UserID          string           `json:"userId" gorm:"uniqueIndex;column:user_id"`
	DisplayName     string           `json:"displayName" gorm:"column:display_name"`
	PictureURL      string           `json:"pictureUrl" gorm:"column:picture_url"`
	StatusMessage   string           `json:"statusMessage" gorm:"column:status_message"`
	Language        string           `json:"language" gorm:"column:language"`
	LineAccessToken string           `json:"line_access_token" gorm:"column:line_access_token"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
