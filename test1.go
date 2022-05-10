package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
)

func main1() {
	client, err := authzed.NewClient(
		"localhost:50051",
		grpcutil.WithInsecureBearerToken("somerandomkeyhere"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("unable to initialize client: %s", err)
	}

	request := &pb.WriteRelationshipsRequest{Updates: []*pb.RelationshipUpdate{
		//{ // Emilia is a Writer on Post 1
		//	Operation: pb.RelationshipUpdate_OPERATION_CREATE,
		//	Relationship: &pb.Relationship{
		//		Resource: &pb.ObjectReference{
		//			ObjectType: "document",
		//			ObjectId:   "1",
		//		},
		//		Relation: "writer",
		//		Subject: &pb.SubjectReference{
		//			Object: &pb.ObjectReference{
		//				ObjectType: "user",
		//				ObjectId:   "emilia",
		//			},
		//		},
		//	},
		//},
		//{ // Beatrice is a Reader on Post 1
		//	Operation: pb.RelationshipUpdate_OPERATION_CREATE,
		//	Relationship: &pb.Relationship{
		//		Resource: &pb.ObjectReference{
		//			ObjectType: "document",
		//			ObjectId:   "1",
		//		},
		//		Relation: "reader",
		//		Subject: &pb.SubjectReference{
		//			Object: &pb.ObjectReference{
		//				ObjectType: "user",
		//				ObjectId:   "beatrice",
		//			},
		//		},
		//	},
		//},
		//{
		//	Operation: pb.RelationshipUpdate_OPERATION_CREATE,
		//	Relationship: &pb.Relationship{
		//		Resource: nil,
		//		Relation: "",
		//		Subject:  nil,
		//	},
		//},
		{ //test
			Operation: pb.RelationshipUpdate_OPERATION_CREATE,
			Relationship: &pb.Relationship{
				Resource: &pb.ObjectReference{
					ObjectType: "organization",
					ObjectId:   "1",
				},
				Relation: "resource",
				Subject: &pb.SubjectReference{
					Object: &pb.ObjectReference{
						ObjectType: "resource",
						ObjectId:   "1",
					},
				},
			},
		},
	}}

	resp, err := client.WriteRelationships(context.Background(), request)
	if err != nil {
		log.Fatalf("failed to write relations: %s", err)
	}
	fmt.Println(resp.WrittenAt.Token)
}
