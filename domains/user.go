package domains

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID         uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	CustomID   string    `gorm:"uniqueIndex" json:"custom_id"`
	Role       string    `json:"role"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Status     string    `json:"status"`
	Parent     string    `json:"parent"`
	Reference  string    `json:"reference"`
	PackageID  string    `json:"package_id"`
	FleetToken string    `json:"fleet_token"`
	LineToken  string    `json:"line_token"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
