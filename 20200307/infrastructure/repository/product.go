package repository

import (
	"context"
	"errors"
	"github.com/rennnosuke/rens-blog-codes/20200307/domain/entity"
	"time"
)

type ProductRepositoryImpl struct{}

func (r *ProductRepositoryImpl) GetProducts(ctx context.Context) ([]entity.Product, error) {

	c := make(chan []entity.Product, 1)
	childCtx, _ := context.WithTimeout(ctx, time.Second*5)

	go func(ctx context.Context) { c <- r.FindProducts(ctx) }(childCtx)

	select {
	case <-childCtx.Done():
		<-c
		return nil, errors.New("query canceled")
	case prods := <-c:
		return prods, nil
	}

}

func (r *ProductRepositoryImpl) FindProducts(ctx context.Context) []entity.Product {
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
	}
}
