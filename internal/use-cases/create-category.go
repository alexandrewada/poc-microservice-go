package use_cases

import (
	"log"

	"github.com/alexandrewada/microservice-go-estudos/internal/entities"
)

type createCategoryUseCase struct {
	//
}

func NewcreateCategoryUseCase() *createCategoryUseCase {
	return &createCategoryUseCase{}
}

func (u *createCategoryUseCase) Execute(name string) error {

	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	log.Println("Category created:", category)

	return nil

}
