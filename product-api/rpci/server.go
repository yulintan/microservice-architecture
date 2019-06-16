package rpci

import (
	"context"
	"strconv"

	pb "github.com/yulintan/microservice-architecture/pb/products"
	"github.com/yulintan/microservice-architecture/product-api/internal/products"
)

type server struct {
	service products.Service
}

func New(service products.Service) pb.ProductServiceServer {
	return &server{
		service: service,
	}
}

func (s *server) GetProductByID(ctx context.Context, req *pb.GetProductByIDRequest) (*pb.GetProductResponse, error) {
	//return dummy data
	product, err := s.service.Get(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.GetProductResponse{
		Product: &pb.Product{
			Id:       req.Id,
			Name:     "Test Name " + strconv.Itoa(int(req.Id)),
			Price:    product.Price,
			Currency: product.Currency,
		},
	}, nil
}
