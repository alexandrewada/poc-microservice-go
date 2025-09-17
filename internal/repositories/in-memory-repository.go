package repositories

import "github.com/alexandrewada/microservice-go-estudos/internal/entities"

type inMemoryCategoryRepository struct {
	db []entities.Category
}

func NewInMemoryCategoryRepository() *inMemoryCategoryRepository {
	return &inMemoryCategoryRepository{
		db: make([]entities.Category, 0),
	}
}

func (r *inMemoryCategoryRepository) Save(category *entities.Category) ([]entities.Category, error) {
	r.db = append(r.db, *category)
	return r.db, nil
}

func (r *inMemoryCategoryRepository) List() ([]entities.Category, error) {
	return r.db, nil
}
