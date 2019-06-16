package products

import (
	"context"

	pb "github.com/yulintan/microservice-architecture/pb/shops"
)

type Service interface {
	Get(context.Context, int) (*Product, error)
}

type productService struct {
	shopRPCClient pb.ShopServiceClient
}

func NewService(shopRPCClient pb.ShopServiceClient) Service {
	return &productService{
		shopRPCClient: shopRPCClient,
	}
}

func (s *productService) Get(ctx context.Context, id int) (*Product, error) {
	// your business logic here, return dummy data
	resp, err := s.shopRPCClient.GetShopByID(ctx, &pb.GetShopByIDRequest{Id: 1})
	if err != nil {
		return nil, err
	}

	return &Product{
		ID:       1,
		Name:     "test product name",
		Price:    "100.00",
		Currency: resp.Shop.Currency,
	}, nil
}
