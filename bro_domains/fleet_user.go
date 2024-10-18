package bro_domains

import (
	"time"

	"github.com/gofrs/uuid"
)

type FleetUser struct {
	OwnerId       uuid.UUID `json:"owner_id" gorm:"primaryKey;column:owner_id"`
	FleetUsername string    `json:"fleet_username" gorm:"uniqueIndex;column:fleet_username"`
	FleetPassword string    `json:"fleet_password" gorm:"column:fleet_password"`
	CompanyID     uint      `json:"company_id" gorm:"column:company_id"`
	StaffID       uint      `json:"staff_id" gorm:"column:staff_id"`
	FleetToken    string    `json:"fleet_token" gorm:"column:fleet_token"`
	IsVerify      bool      `json:"is_verify" gorm:"column:is_verify"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
