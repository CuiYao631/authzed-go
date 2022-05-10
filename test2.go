package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
)

func main2() {
	client, err := authzed.NewClient(
		"localhost:50051",
		grpcutil.WithInsecureBearerToken("somerandomkeyhere"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("unable to initialize client: %s", err)
	}

	ctx := context.Background()

	emilia := &pb.SubjectReference{Object: &pb.ObjectReference{
		ObjectType: "resource",
		ObjectId:   "1",
	}}

	firstPost := &pb.ObjectReference{
		ObjectType: "organization",
		ObjectId:   "2",
	}

	resp, err := client.CheckPermission(ctx, &pb.CheckPermissionRequest{
		Resource:   firstPost,
		Permission: "resource",
		Subject:    emilia,
	})
	if err != nil {
		log.Fatalf("failed to check permission: %s", err)
	}
	log.Println(resp)
	//b := &pb.RelationshipFilter{
	//	ResourceType:       "usergroup",
	//	OptionalResourceId: "1",
	//	OptionalRelation:   "manager",
	//	OptionalSubjectFilter: &pb.SubjectFilter{
	//		SubjectType: "user",
	//		//OptionalSubjectId: "11",
	//	},
	//}
	//res, err := client.ReadRelationships(ctx, &pb.ReadRelationshipsRequest{
	//	RelationshipFilter: b,
	//})
	//
	//ress, err := res.Recv()
	//
	//log.Println(ress)

	//res, err := client.ExpandPermissionTree(ctx, &pb.ExpandPermissionTreeRequest{
	//	Resource: &pb.ObjectReference{
	//		ObjectType: "resource",
	//		ObjectId:   "1",
	//	},
	//	Permission: "manager",
	//})
	//log.Println(res.TreeRoot.GetLeaf().GetSubjects())
	//res, err := client.LookupResources(ctx, &pb.LookupResourcesRequest{
	//	ResourceObjectType: "resource",
	//	Permission:         "manager",
	//	Subject: &pb.SubjectReference{Object: &pb.ObjectReference{
	//		ObjectType: "user",
	//		ObjectId:   "21",
	//	}},
	//})
	//ress, err := res.Recv()
	//
	//log.Println(ress.GetResourceObjectId())

}
