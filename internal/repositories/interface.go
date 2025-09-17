package repositories

import "github.com/alexandrewada/microservice-go-estudos/internal/entities"

type ICategoryRepository interface {
	Save(category *entities.Category) ([]entities.Category, error)
	List() ([]entities.Category, error)
}
