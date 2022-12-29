package service

import (
	"errors"
	"unit-test/entity"
	"unit-test/repository"
)

type CategoryService struct {
	Respository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Respository.FindById(id)
	if category == nil {
		return category, errors.New("Cetogry Not Found")
	} else {
		return category, nil
	}
}
