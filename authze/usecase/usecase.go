package usecase

import (
	"authzed-go/authze/entity"
	"context"
	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
)

// API https://buf.build/authzed/api/docs/main:authzed.api.v1#authzed.api.v1.PermissionsService.WriteRelationships

type Usecase interface {
	//添加关系
	AddPerm(ctx context.Context, objectType, objectId, subjectType, subjectId, relation string) (string, error)
	//读取关系
	ReadPerm(ctx context.Context, resourceType string) ([]entity.PermResponse, error)
	//检查关系
	CheckPerm(ctx context.Context, objectType, objectId, subjectType, subjectId, permission string) string
	//更新关系
	UpdatePerm(ctx context.Context)
	//返回所有Resources
	LookupResources(ctx context.Context, objectType, permission, subjectType, subjectId string) ([]entity.ResourcesResponse, error)
	//返回所有Subjects
	LookupSubjects(ctx context.Context, objectType, objectId, permission, subjectType string) ([]entity.SubjectResponse, error)
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

func (u usecase) AddPerm(ctx context.Context, objectType, objectId, subjectType, subjectId, relation string) (string, error) {
	request := &pb.WriteRelationshipsRequest{Updates: []*pb.RelationshipUpdate{
		{
			Operation: pb.RelationshipUpdate_OPERATION_CREATE,
			Relationship: &pb.Relationship{
				Resource: &pb.ObjectReference{
					ObjectType: objectType,
					ObjectId:   objectId,
				},
				Relation: relation,
				Subject: &pb.SubjectReference{
					Object: &pb.ObjectReference{
						ObjectType: subjectType,
						ObjectId:   subjectId,
					},
				},
			},
		},
	}}
	resp, err := u.client.WriteRelationships(ctx, request)
	if err != nil {
		return "", err
	}
	return resp.WrittenAt.Token, nil
}
func (u usecase) ReadPerm(ctx context.Context, resourceType string) ([]entity.PermResponse, error) {
	var relation []entity.PermResponse
	req := &pb.ReadRelationshipsRequest{
		RelationshipFilter: &pb.RelationshipFilter{
			ResourceType: resourceType,
			//OptionalSubjectFilter: &pb.SubjectFilter{
			//	SubjectType:       "user",
			//	OptionalSubjectId: "123",
			//},
		},
	}
	resp, err := u.client.ReadRelationships(ctx, req)
	if err != nil {
		log.Fatalf("failed to write schema: %s", err)
	}

	for {
		feature, err := resp.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		relation = append(relation, entity.PermResponse{
			Token: feature.ReadAt.Token,
			Relationship: struct {
				ObjectType  string
				ObjectId    string
				Relation    string
				SubjectType string
				SubjectId   string
			}{
				ObjectType:  feature.Relationship.Resource.ObjectType,
				ObjectId:    feature.Relationship.Resource.ObjectId,
				Relation:    feature.Relationship.Relation,
				SubjectType: feature.Relationship.Subject.Object.ObjectType,
				SubjectId:   feature.Relationship.Subject.Object.ObjectId,
			},
		})
		log.Println(feature)
	}
	return relation, nil
}
func (u usecase) CheckPerm(ctx context.Context, objectType, objectId, subjectType, subjectId, permission string) string {
	Subject := &pb.SubjectReference{Object: &pb.ObjectReference{
		ObjectType: subjectType,
		ObjectId:   subjectId,
	}}

	Resource := &pb.ObjectReference{
		ObjectType: objectType,
		ObjectId:   objectId,
	}

	resp, err := u.client.CheckPermission(ctx, &pb.CheckPermissionRequest{
		Resource:   Resource,
		Permission: permission,
		Subject:    Subject,
	})
	if err != nil {
		log.Fatalf("failed to check permission: %s", err)
	}
	str := ""
	if resp.Permissionship == pb.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION {
		str = "允许"
	} else if resp.Permissionship == pb.CheckPermissionResponse_PERMISSIONSHIP_NO_PERMISSION {
		str = "拒绝"
	} else {
		str = "未知"
	}

	return str
}
func (u usecase) LookupResources(ctx context.Context, objectType, permission, subjectType, subjectId string) ([]entity.ResourcesResponse, error) {

	req := &pb.LookupResourcesRequest{
		ResourceObjectType: objectType,
		Permission:         permission,
		Subject: &pb.SubjectReference{
			Object: &pb.ObjectReference{
				ObjectType: subjectType,
				ObjectId:   subjectId,
			},
		},
	}
	resp, err := u.client.LookupResources(ctx, req)
	if err != nil {
		return nil, err
	}
	var resResp []entity.ResourcesResponse
	for {
		feature, err := resp.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		resResp = append(resResp, entity.ResourcesResponse{
			Token:    feature.LookedUpAt.Token,
			ObjectId: feature.ResourceObjectId,
		})
	}
	return resResp, nil
}
func (u usecase) LookupSubjects(ctx context.Context, objectType, objectId, permission, subjectType string) ([]entity.SubjectResponse, error) {
	req := &pb.LookupSubjectsRequest{
		Resource: &pb.ObjectReference{
			ObjectType: objectType,
			ObjectId:   objectId,
		},
		Permission:        permission,
		SubjectObjectType: subjectType,
	}
	resp, err := u.client.LookupSubjects(ctx, req)
	if err != nil {
		log.Fatalf("failed to check permission: %s", err)
	}
	var resResp []entity.SubjectResponse
	for {
		ress, err := resp.Recv()
		if err == io.EOF {
			break
		}
		resResp = append(resResp, entity.SubjectResponse{
			Token:     ress.LookedUpAt.Token,
			SubjectID: ress.SubjectObjectId,
		})
	}
	return resResp, err
}

func (u usecase) UpdatePerm(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (u usecase) DelPerm(ctx context.Context) {
	//req:=&pb.WriteRelationshipsRequest{
	//	Updates:,
	//	OptionalPreconditions: nil,
	//}
	panic("implement me")
}
func (u usecase) ReadSchema(ctx context.Context) (string, error) {
	req := &pb.ReadSchemaRequest{}
	resp, err := u.client.ReadSchema(context.Background(), req)
	if err != nil {
		return "", status.Error(codes.InvalidArgument, "failed to Read schema"+err.Error())
		//log.Fatalf("failed to write schema: %s", err)
	}
	log.Println("ReadSchema")
	//log.Println(resp.SchemaText)
	return resp.SchemaText, nil

}
func (u usecase) WriteScheme(ctx context.Context, scheme string) error {
	request := &pb.WriteSchemaRequest{Schema: scheme}
	res, er := u.client.WriteSchema(context.Background(), request)
	if er != nil {
		return status.Error(codes.InvalidArgument, "failed to write schema"+er.Error())
		//log.Fatalf("failed to write schema: %s", err)
	}
	log.Println("WriteScheme")
	log.Println(res)
	return nil
}
