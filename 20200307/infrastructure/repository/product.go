package repository

import (
	"context"
	"github.com/rennnosuke/rens-blog-codes/20200307/domain/entity"
)

type ProductRepositoryImpl struct{}

func (r *ProductRepositoryImpl) GetProducts(ctx context.Context) ([]entity.Product, error) {
	return []entity.Product{
		{
			Name:  "Chocolate",
			Price: 100,
		},
		{
			Name:  "Orange Juice",
			Price: 150,
		},
		{
			Name:  "Candy",
			Price: 100,
		},
	}, nil
}
