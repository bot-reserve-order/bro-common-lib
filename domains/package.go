package domains

type Package struct {
	ID string `gorm:"primaryKey" json:"id"`
}