package main

import (
	"authzed-go/authze"
	"github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"os"
)

const (
	PORT = ":9192"
)

func main() {
	client, err := authzed.NewClient(
		os.Getenv("HOST"),
		// NOTE: For SpiceDB behind TLS, use:
		// grpcutil.WithBearerToken("somerandomkeyhere"),
		// grpcutil.WithSystemCerts(grpcutil.VerifyCA),
		grpcutil.WithInsecureBearerToken("somekey"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("unable to initialize client: %s", err)
	}
	log.Println(os.Getenv("HOST"))
	s := grpc.NewServer()

	auth := authze.NewAuthzed(client)
	auth.InitServer(s)
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("faled to listen: %v", err)
	}
	//start server
	log.Println("starting grpc server, listening on", PORT)
	if err := s.Serve(lis); err != nil {
		log.Fatal("start grpc server failed with", err)
	}

}
