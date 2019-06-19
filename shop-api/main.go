package main

import (
	"flag"
	"net"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/yulintan/microservice-architecture/lib/grpclib"
	pb "github.com/yulintan/microservice-architecture/pb/shops"
	"github.com/yulintan/microservice-architecture/shop-api/internal/shops"
	"github.com/yulintan/microservice-architecture/shop-api/rpci"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	echoEndpoint = flag.String("echo_endpoint", "localhost:9090", "endpoint of YourService")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	s := grpc.NewServer()

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	shopService := shops.NewService()
	pb.RegisterShopServiceServer(s, rpci.New(shopService))
	go s.Serve(lis)

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}),
		runtime.WithProtoErrorHandler(grpclib.ErrorHandler),
	)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = pb.RegisterShopServiceHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
