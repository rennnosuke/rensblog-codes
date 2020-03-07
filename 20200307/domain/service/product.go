package service

import (
	"context"
	"github.com/rennnosuke/rens-blog-codes/20200307/domain/entity"
	"github.com/rennnosuke/rens-blog-codes/20200307/domain/repository"
)

type ProductService struct {
	Repo repository.ProductRepository
}

func (s *ProductService) GetProducts(ctx context.Context) ([]entity.Product, error) {
	return s.Repo.GetProducts(ctx)
}
