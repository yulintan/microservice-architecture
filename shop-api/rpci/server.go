package rpci

import (
	"context"

	"github.com/yulintan/microservice-architecture/lib/grpclib"
	pb "github.com/yulintan/microservice-architecture/pb/shops"
	"github.com/yulintan/microservice-architecture/shop-api/internal/shops"
)

type server struct {
	shopService shops.Service
}

func New(ss shops.Service) pb.ShopServiceServer {
	return &server{
		shopService: ss,
	}
}

func (s *server) GetShopCount(ctx context.Context, req *pb.GetShopCountRequest) (*pb.GetShopCountResponse, error) {
	count, err := s.shopService.GetTotalShops(ctx)
	if err != nil {
		return nil, grpclib.Error(err)
	}

	return &pb.GetShopCountResponse{
		Count: int32(count),
	}, nil
}

func (s *server) GetShopByID(ctx context.Context, req *pb.GetShopByIDRequest) (*pb.GetShopResponse, error) {
	shop, err := s.shopService.GetById(ctx, int(req.Id))
	if err != nil {
		return nil, grpclib.Error(err)
	}

	return &pb.GetShopResponse{
		Shop: &pb.Shop{
			Id:         int64(shop.ID),
			ShopDomain: shop.ShopDomain,
			Currency:   shop.Currency,
		},
	}, nil
}
