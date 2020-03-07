package main

import (
	"context"
	"fmt"
	"github.com/rennnosuke/rens-blog-codes/20200307/domain/service"
	"github.com/rennnosuke/rens-blog-codes/20200307/infrastructure/repository"
)

func main() {
	r := repository.ProductRepositoryImpl{}
	s := service.ProductService{Repo: &r}

	ctx := context.Background()

	prods, err := s.GetProducts(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(prods)
}
