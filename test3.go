package main

import (
	"context"
	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main3() {
	client, err := authzed.NewClient(
		"localhost:50051",
		grpcutil.WithInsecureBearerToken("somerandomkeyhere"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("unable to initialize client: %s", err)
	}

	req := &pb.ReadSchemaRequest{}
	resp, err := client.ReadSchema(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to write schema: %s", err)
	}
	log.Println(resp.SchemaText)

	//newSchema := resp.SchemaText + "definition user2 {}"
	//
	//request := &pb.WriteSchemaRequest{Schema: newSchema}
	//res, er := client.WriteSchema(context.Background(), request)
	//if er != nil {
	//	log.Fatalf("failed to write schema: %s", err)
	//}
	//log.Println(res)
	//r := &pb.ReadSchemaRequest{}
	//re, err := client.ReadSchema(context.Background(), r)
	//if err != nil {
	//	log.Fatalf("failed to write schema: %s", err)
	//}
	//log.Println(re.SchemaText)
}
