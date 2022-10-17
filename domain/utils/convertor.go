package utils

import (
	"faceit/domain/dto"
	"faceit/domain/entity"
)

func UserDTOFromEntity(entity *entity.User) *dto.User {
	return &dto.User{
		ID:        entity.ID,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		NickName:  entity.NickName,
		Email:     entity.Email,
		Country:   entity.Country,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

func UserEntityFromDTO(dto *dto.User) *entity.User {
	return &entity.User{
		ID:        dto.ID,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		NickName:  dto.NickName,
		Email:     dto.Email,
		Country:   dto.Country,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}
