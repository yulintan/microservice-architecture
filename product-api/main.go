package main

import (
	"context"
	"flag"
	"net"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/yulintan/microservice-architecture/lib/grpclib"
	pbproduct "github.com/yulintan/microservice-architecture/pb/products"
	pbshop "github.com/yulintan/microservice-architecture/pb/shops"
	"github.com/yulintan/microservice-architecture/product-api/internal/products"
	"github.com/yulintan/microservice-architecture/product-api/rpci"
	"google.golang.org/grpc"
)

var (
	echoEndpoint   = flag.String("echo_endpoint", "localhost:9091", "endpoint of YourService")
	shopRPCAddress = "localhost:9090"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	s := grpc.NewServer()

	lis, err := net.Listen("tcp", ":9091")
	if err != nil {
		panic(err)
	}

	conn, err := grpc.Dial(shopRPCAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	shopRPCClient := pbshop.NewShopServiceClient(conn)
	productService := products.NewService(shopRPCClient)

	pbproduct.RegisterProductServiceServer(s, rpci.New(productService))
	go s.Serve(lis)

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}),
		runtime.WithProtoErrorHandler(grpclib.ErrorHandler),
	)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = pbproduct.RegisterProductServiceHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
