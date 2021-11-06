package service

import (
	"galaxy/api/repository"
	"galaxy/models"
)

type ConverterService struct {
	repository repository.ConverterRepository
}

func NewConverterService(r repository.ConverterRepository) ConverterService {
	return ConverterService{
		repository: r,
	}
}

func (p ConverterService) ConvertRoman(converter models.Converter) (int, error) {
	return p.repository.ConvertRoman(converter)
}
