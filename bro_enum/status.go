package bro_enum

type Status string

const (
	Active  Status = "active"
	Pending Status = "pending"
	Disable Status = "disable"
)