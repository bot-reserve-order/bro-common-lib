package bro_domains

import (
	"time"

	"github.com/bot-reserve-order/bro-common-lib/bro_enum"
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

type WithdrawRequest struct {
	ID            uint            `gorm:"primaryKey" json:"id"`
	Note          string          `json:"note"`
	Status        bro_enum.Status `json:"status"`
	Amount        decimal.Decimal `json:"amount"`
	FounderUserID uuid.UUID       `json:"founder_user_id"`
	CreditID      uint            `json:"credit_id"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
