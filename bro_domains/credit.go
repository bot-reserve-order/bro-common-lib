package bro_domains

import (
	"time"

	"github.com/bot-reserve-order/bro-common-lib/bro_enum"
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

type Credit struct {
	ID          uint            `gorm:"primaryKey" json:"id"`
	UserID      uuid.UUID       `gorm:"type:uuid;index" json:"user_id" grom:"column:user_id"`
	Amount      decimal.Decimal `json:"amount" grom:"column:amount"`
	Note        string          `json:"note" grom:"column:note"`
	Status      bro_enum.Status `json:"status" grom:"column:status"`
	ExternalRef string          `json:"external_ref" grom:"column:external_ref;index"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TrueWalletCallbackObj struct {
	EventType    string          `json:"event_type"`
	ReceivedTime string          `json:"received_time"`
	Amount       decimal.Decimal `json:"amount"`
	SenderMobile string          `json:"sender_mobile"`
	Message      string          `json:"message"`
	Iat          uint            `json:"iat"`
}
