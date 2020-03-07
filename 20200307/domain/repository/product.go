package repository

import (
	"context"
	"github.com/rennnosuke/rens-blog-codes/20200307/domain/entity"
)

type ProductRepository interface {
	GetProducts(ctx context.Context) ([]entity.Product, error)
}
