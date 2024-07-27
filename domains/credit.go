package domains

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

type Credit struct {
	ID          uint            `gorm:"primaryKey" json:"id"`
	UserID      uuid.UUID       `gorm:"type:uuid;" json:"user_id"`
	Amount      decimal.Decimal `json:"amount"`
	Note        string          `json:"note"`
	Status      string          `json:"status"`
	ExternalRef string          `gorm:"index" json:"external_ref"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
