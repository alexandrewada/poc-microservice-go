package use_cases

import (
	"log"

	"github.com/alexandrewada/microservice-go-estudos/internal/entities"
	"github.com/alexandrewada/microservice-go-estudos/internal/repositories"
)

type listCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewlistCategoryUseCase(repository repositories.ICategoryRepository) *listCategoryUseCase {
	return &listCategoryUseCase{repository}
}

func (u *listCategoryUseCase) Execute(name string) error {

	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	categories, err := u.repository.Save(category)

	if err != nil {
		log.Printf("Error saving category: %v", err)
		return err
	}

	log.Printf("Categories: %+v", categories)

	return nil

}
