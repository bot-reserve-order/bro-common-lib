package bro_domains

import "github.com/gofrs/uuid"

type FleetSocketSession struct {
	ID           uuid.UUID `json:"id"`
	FmsSessionID string    `json:"fms_session_id"`
	CompanyID    uint      `json:"company_id" gorm:"column:company_id"`
	StaffID      uint      `json:"staff_id" gorm:"column:staff_id"`
}
