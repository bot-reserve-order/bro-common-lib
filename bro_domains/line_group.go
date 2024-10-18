package bro_domains

import "github.com/gofrs/uuid"

type LineGroup struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();uniqueIndex" json:"id"`
	OwnerId     uuid.UUID `json:"owner_id" gorm:"type:uuid;column:owner_id"`
	GroupId     string    `json:"groupId" gorm:"column:group_id"`
	GroupName   string    `json:"groupName" gorm:"column:group_name"`
	MemberCount uint      `json:"member_count" gorm:"column:member_count"`
	PictureURL  string    `json:"pictureUrl" gorm:"column:picture_url"`
}
