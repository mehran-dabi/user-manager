package entity

import (
	"faceit/domain/user/dto"
)

type Filter struct {
	Country  string
	NickName string
}

func FilterEntityFromDTO(dto *dto.Filter) *Filter {
	return &Filter{
		Country:  dto.Country,
		NickName: dto.NickName,
	}
}
