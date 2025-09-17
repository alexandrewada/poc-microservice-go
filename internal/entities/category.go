package entities

import (
	"fmt"
	"math"
	"math/rand/v2"
	"time"
)

type Category struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCategory(name string) (*Category, error) {

	category := &Category{
		ID:        uint(math.Round(rand.Float64() * 1000)), // Simulate ID generation
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := category.IsValid()

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (c *Category) IsValid() error {
	if c.Name == "" {
		return fmt.Errorf("name is required")
	}

	return nil
}
