package usecase

import (
	"context"
	"fmt"
	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type Usecase interface {
	AddPerm(ctx context.Context)
	UpdatePerm(ctx context.Context)
	DelPerm(ctx context.Context)
	ReadSchema(ctx context.Context) (string, error)
	WriteScheme(ctx context.Context, scheme string) error
}
type usecase struct {
	client *authzed.Client
}

func NewUsecase(client *authzed.Client) *usecase {
	return &usecase{client: client}
}

func (u usecase) AddPerm(ctx context.Context) {
	request := &pb.WriteRelationshipsRequest{Updates: []*pb.RelationshipUpdate{
		{
			Operation: pb.RelationshipUpdate_OPERATION_CREATE,
			Relationship: &pb.Relationship{
				Resource: &pb.ObjectReference{
					ObjectType: "document",
					ObjectId:   "1",
				},
				Relation: "writer",
				Subject: &pb.SubjectReference{
					Object: &pb.ObjectReference{
						ObjectType: "user",
						ObjectId:   "emilia",
					},
				},
			},
		},
	}}
	resp, err := u.client.WriteRelationships(context.Background(), request)
	if err != nil {
		log.Fatalf("failed to write relations: %s", err)
	}
	fmt.Println(resp.WrittenAt.Token)
}

func (u usecase) UpdatePerm(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (u usecase) DelPerm(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}
func (u usecase) ReadSchema(ctx context.Context) (string, error) {
	req := &pb.ReadSchemaRequest{}
	resp, err := u.client.ReadSchema(context.Background(), req)
	if err != nil {
		return "", status.Error(codes.InvalidArgument, "failed to write schema")
		//log.Fatalf("failed to write schema: %s", err)
	}
	//log.Println(resp.SchemaText)
	return resp.SchemaText, nil

}
func (u usecase) WriteScheme(ctx context.Context, scheme string) error {
	request := &pb.WriteSchemaRequest{Schema: scheme}
	res, er := u.client.WriteSchema(context.Background(), request)
	if er != nil {
		return status.Error(codes.InvalidArgument, "failed to write schema")
		//log.Fatalf("failed to write schema: %s", err)
	}
	log.Println(res)
	return nil
}
