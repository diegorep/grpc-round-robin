package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"


	pb "./services/threes"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	port       = flag.Int("port", 50049, "The server port")
)

type threesServer struct {
}

func (s *threesServer) EchoThree(ctx context.Context, req *pb.GetThreeRequest) (*pb.GetThreeResponse, error) {
    fmt.Println("serving echo request")
    retVal := float32(-3.0)
    if req.Value.Value == 3 {
        retVal = float32(3.0)
    }
    message := &pb.Three{Value: retVal}
    return &pb.GetThreeResponse{Value: message}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = testdata.Path("server1.pem")
		}
		if *keyFile == "" {
			*keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterThreesServer(grpcServer, &threesServer{})
    fmt.Println("Server listening on port", *port)
	grpcServer.Serve(lis)
}
