package service

import (
	"galaxy/models"
	"server/api/repository"
	"server/models"
)

type ConverterService struct {
	repository repository.ConverterRepository
}

func NewPostService(r repository.ConverterRepository) ConverterService {
	return ConverterService{
		repository: r,
	}
}

func (p ConverterService) Find(converter models.Converter) (models.Converter, error) {
	return p.repository.Find(converter)
}
