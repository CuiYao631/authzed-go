package main

import (
	"authzed-go/authze"
	"github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
)

const schema = `

definition user {}

definition resource {
    relation manager: user | usergroup#member | usergroup#manager
    relation viewer: user | usergroup#member | usergroup#manager

    permission manage = manager
    permission view = viewer + manager
}

definition usergroup {
    relation manager: user | usergroup#member | usergroup#manager
    relation direct_member: user | usergroup#member | usergroup#manager
	
	relation resource: resource 

    permission member = direct_member + manager
}

definition organization {
    relation group: usergroup
    relation administrator: user | usergroup#member | usergroup#manager
    relation direct_member: user

    relation resource: resource

    permission admin = administrator
    permission member = direct_member + administrator + group->member
}
`
const (
	PORT = ":9192"
)

func main() {
	client, err := authzed.NewClient(
		"localhost:50051",
		// NOTE: For SpiceDB behind TLS, use:
		// grpcutil.WithBearerToken("somerandomkeyhere"),
		// grpcutil.WithSystemCerts(grpcutil.VerifyCA),
		grpcutil.WithInsecureBearerToken("somerandomkeyhere"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("unable to initialize client: %s", err)
	}

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
