package entity

import "time"

type Filter struct {
	Country   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
