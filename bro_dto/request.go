package bro_dto

import (
	"time"

	"github.com/bot-reserve-order/bro-common-lib/bro_enum"
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

type GetConditionsRequest struct {
	Status          []bro_enum.Status  `json:"status" form:"status"`
	Area            []string  `json:"area" form:"area"`
	Cartype         []string  `json:"car_type" form:"car_type"`
	TrackStart      string    `json:"track_start" form:"track_start"`
	TrackEnd        string    `json:"track_end" form:"track_end"`
	MaxRange        *uint     `json:"max_range" form:"max_range"`
	SerialID        *string    `json:"serial_id" form:"serial_id"`
	UserID          uuid.UUID `json:"user_id" form:"user_id"`
	UserRootID      uuid.UUID `json:"user_root_id" form:"user_root_id"`
	FounderUserID   uuid.UUID `json:"founder_user_id" form:"founder_user_id"`
	ReferenceUserID uuid.UUID `json:"reference_user_id" form:"reference_user_id"`
	CreatedAtStart  time.Time `json:"created_at_start" form:"created_at_start"`
	CreatedAtEnd    time.Time `json:"created_at_end" form:"created_at_end"`
	UpdatedAtStart  time.Time `json:"updated_at_start" form:"updated_at_start"`
	UpdatedAtEnd    time.Time `json:"updated_at_end" form:"updated_at_end"`
	PerPage         uint      `json:"per_page" form:"per_page"`
	PageNum         uint      `json:"page_num" form:"page_num"`
}

// subuser login and founder login
type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RootLoginRegister struct {
	Login
	Reference string `json:"reference" validate:"required"`
	LineToken string `json:"line_token" validate:"required"`
}

type FounderRegister struct {
	Login
	TrueWallet string `json:"true_wallet" validate:"required,number"`
}

type SubUserRegister struct {
	Login
	LineToken string `json:"line_token"`
}

type UpdateUserPassword struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

type UpdateUserDisplayName struct {
	NewDisplayName string `json:"new_display_name" validate:"required"`
}

type UpdateUserCustomID struct {
	NewCustomID string `json:"new_custom_id" validate:"required"`
}

type UpdateUserStatus struct {
	NewStatus string `json:"new_status" validate:"required"`
}

type UpdateLineToken struct {
	NewLineToken string `json:"new_line_token" validate:"required"`
}

type BoostCondition struct {
	Boost decimal.Decimal `json:"boost" validate:"required"`
}

type TrueWalletCallback struct {
	Message string `json:"message"`
}

type GetUsersRequest struct {
	ID          []uuid.UUID         `json:"ids"`
	CustomID    []string            `json:"custom_ids"`
	DisplayName string              `json:"display_name"`
	Role        []bro_enum.RoleUser `json:"roles"`
	Username    []string            `json:"usernames"`
	Status      []bro_enum.Status   `json:"status"`
	Parent      []uuid.UUID         `json:"parents"`
	CompanyID   []uint              `json:"company_ids"`
	StaffID     []uint              `json:"staff_ids"`
	Reference   []string            `json:"references"`
	PackageID   []bro_enum.Package  `json:"package_ids"`
	FleetToken  []string            `json:"fleet_tokens"`
	LineToken   []string            `json:"line_tokens"`
	TrueWallet  []string            `json:"true_wallets"`
	PerPage     uint                `json:"per_page" form:"per_page"`
	PageNum     uint                `json:"page_num" form:"page_num"`
}
