package service

import (
	"context"

	github.com/igorlev91/golang-grpc-ecommerce/pb/cart"
)

type ICartService interface {
	AddProductToCart(ctx context.Context, request *cart.AddProductToCartRequest) (*cart.AddProductToCartResponse, error)
}
