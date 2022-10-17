package entity

import (
	"faceit/domain/dto"
	"time"
)

type Filter struct {
	Country   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FilterEntityFromDTO(dto *dto.Filter) *Filter {
	return &Filter{
		Country:   dto.Country,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}
