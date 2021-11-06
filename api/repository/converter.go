package repository

import (
	"errors"
	"galaxy/base"
	"galaxy/models"
	"strings"
)

type ConverterRepository struct {
	db base.Database
}

func NewConverterRepository(db base.Database) ConverterRepository {
	return ConverterRepository{
		db: db,
	}
}

func GetValue(x string) int {
	var val int
	switch {

	case x == "glob":
		val = 1
	case x == "prok":
		val = 5
	case x == "pish":
		val = 10
	case x == "tegj":
		val = 50
	default:
		val = 0

	}

	return val
}

func (p ConverterRepository) ConvertRoman(converter models.Converter) (int, error) {

	var result int

	data := strings.Fields(converter.Number)

	for i := 0; i < len(data); i++ {

		p1 := GetValue(string(data[i]))
		if p1 == 0 {
			err := errors.New("error")
			return 0, err
		}

		if i+1 < len(data) {

			p2 := GetValue(string(data[i+1]))
			if p2 == 0 {
				err := errors.New("error")
				return 0, err
			}

			if p1 >= p2 {
				result = result + p1

			} else {
				result = result + p2 - p1

				i++
			}

		} else {
			result = result + p1
		}
	}

	return result, nil
}
