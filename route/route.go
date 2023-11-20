package route

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"

	log "github.com/sirupsen/logrus"

	"main/proto"
	"main/proto/protogen"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	GRPC_PORT  = 9988
	PROXY_PORT = 8888
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", fmt.Sprintf("localhost:%v", GRPC_PORT), "gRPC server endpoint")
)

func RunProxy() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := protogen.RegisterYourServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts) // 1. Replace here when make new protobuf file
	if err != nil {
		return err
	}

	// Refer: https://grpc-ecosystem.github.io/grpc-gateway/docs/operations/inject_router/#adding-custom-routes-to-the-mux
	err = mux.HandlePath("POST", "/sample_grpc", func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
		// Implement your customize route here
	})
	if err != nil {
		return err
	}

	log.Infof("Proxy Server is running at port %v...", PROXY_PORT)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(fmt.Sprintf(":%v", PROXY_PORT), mux)
}

func RunGrpc() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", GRPC_PORT))
	if err != nil {
		log.Errorf("err while create listen %v", err)
		return err
	}

	s := grpc.NewServer()
	protogen.RegisterYourServiceServer(s, proto.YourServiceServer{}) // 2. Replace here when make new protobuf file

	log.Infof("gRPC Gateway is running at port %v...", GRPC_PORT)
	return s.Serve(lis)
}
