package shops

import (
	"context"
)

//go:generate counterfeiter -o fakes/fake_service.go . Service
type Service interface {
	GetById(context.Context, int) (*Shop, error)
	GetTotalShops(ctx context.Context) (int, error)
}

type shopService struct {
}

func NewService() Service {
	return &shopService{}
}

func (s *shopService) GetById(ctx context.Context, id int) (*Shop, error) {
	// your business logic here, return dummy data
	return &Shop{
		ID:         1,
		ShopDomain: "test1.shop.com",
		Currency:   "$",
	}, nil
}

func (s *shopService) GetTotalShops(ctx context.Context) (int, error) {
	// your business logic here, return dummy data
	return 2, nil
}
