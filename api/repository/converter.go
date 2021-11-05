package repository

import (
	"server/base"
	"server/models"
)

type ConverterRepository struct {
	db base.Database
}

func NewConverterRepository(db base.Database) ConverterRepository {
	return ConverterRepository{
		db: db,
	}
}

//Find -> Method for fetching post by id
func (p ConverterRepository) Find(converter models.Converter) (models.Converter, error) {
	var converters models.Converter
	err := p.db.DB.
		Debug().
		Model(&models.Converter{}).
		Where(&converter).
		Take(&converters).Error
	return converters, err
}
