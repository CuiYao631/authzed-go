package controller

import (
	"authzed-go/authze/pb"
	"authzed-go/authze/usecase"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const schema = `

// 部门
definition department {
	relation administrator: user | department#member | department#manager | file#reader | file#writer
	relation manager: user | department#member | department#manager | file#reader | file#writer
	relation head: user | department#member | department#manager | file#reader | file#writer
	relation deputyhead: user | department#member | department#manager | file#reader | file#writer
	relation direct_member: user | department#member | department#manager | file#reader | file#writer
	relation department: department
	permission department_admin = administrator
	permission department_manger = administrator + manager
	permission department_head = administrator + manager + head
	permission department_deputyhead = administrator + manager + head + deputyhead
	permission member = administrator + manager + head + deputyhead + direct_member
	permission view_file = department_manger
}

// 文件
definition file {
	relation department: department

	// 这里利用reader关联到部门，还有用户等，大的权限定义是公司组织架构，文件服务只是一个单独服务的权限定义，
	relation reader: user | department | department#member | department#manager
	relation writer: user | department | department#member | department#manager
	permission view = reader + writer + department->view_file
}

// 用户
definition user {}
`

type Server interface {
	pb.AuthzedSVCServer
}
type server struct {
	uc usecase.Usecase
	pb.UnimplementedAuthzedSVCServer
}

func NewServer(uc usecase.Usecase) Server {
	return &server{
		uc,
		pb.UnimplementedAuthzedSVCServer{},
	}
}

// AddPerm 添加权限
func (s *server) AddPerm(ctx context.Context, req *pb.AddPermRequest) (*pb.AddPermResponse, error) {

	res, err := s.uc.AddPerm(ctx, req.ObjectType, req.ObjectId, req.SubjectType, req.SubjectId, req.Relation)
	if err != nil {
		return nil, err
	}
	return &pb.AddPermResponse{
		ZedToken: res,
	}, nil
}

func (s *server) ReadPerm(ctx context.Context, req *pb.ReadPermRequest) (*pb.ReadPermResponse, error) {
	res, err := s.uc.ReadPerm(ctx, req.ResourceType)
	if err != nil {
		return nil, err
	}
	resp := make([]*pb.PermResponse, 0, len(res))
	for _, v := range res {
		resp = append(resp, &pb.PermResponse{
			Token: v.Token,
			Relationship: &pb.Relationship{
				ObjectType:  v.Relationship.ObjectType,
				ObjectId:    v.Relationship.ObjectId,
				Relation:    v.Relationship.Relation,
				SubjectType: v.Relationship.SubjectType,
				SubjectId:   v.Relationship.SubjectId,
			},
		})
	}
	return &pb.ReadPermResponse{
		PermResponse: resp,
	}, nil
}
func (s *server) CheckPerm(ctx context.Context, req *pb.CheckPermRequest) (*pb.CheckPermResponse, error) {
	res := s.uc.CheckPerm(ctx, req.ObjectType, req.ObjectId, req.SubjectType, req.SubjectId, req.Permission)
	//s.uc.LookupSubjects(ctx)
	return &pb.CheckPermResponse{
		Checked: res,
	}, nil
}

// UpdatePerm 更新权限
func (s *server) UpdatePerm(ctx context.Context, req *pb.UpdatePermRequest) (*pb.UpdatePermResponse, error) {

	return &pb.UpdatePermResponse{}, nil
}

// DelPerm 删除权限
func (s *server) DelPerm(ctx context.Context, req *pb.DelPermRequest) (*pb.DelPermResponse, error) {

	return &pb.DelPermResponse{}, nil
}
func (s *server) LookupResources(ctx context.Context, req *pb.LookupResourcesRequest) (*pb.LookupResourcesResponse, error) {
	res, err := s.uc.LookupResources(ctx, req.ObjectType, req.Permission, req.SubjectType, req.SubjectId)
	if err != nil {
		return nil, err
	}
	resResp := make([]*pb.ResourcesResponse, 0, len(res))
	for _, v := range res {
		resResp = append(resResp, &pb.ResourcesResponse{
			Token:    v.Token,
			ObjectId: v.ObjectId,
		})
	}
	return &pb.LookupResourcesResponse{ResourcesResponse: resResp}, nil
}
func (s *server) LookupSubjects(ctx context.Context, req *pb.LookupSubjectsRequest) (*pb.LookupSubjectsResponse, error) {

	res, err := s.uc.LookupSubjects(ctx, req.ObjectType, req.ObjectId, req.Permission, req.SubjectType)
	if err != nil {
		return nil, err
	}
	resResp := make([]*pb.SubjectsResponse, 0, len(res))
	for _, v := range res {
		resResp = append(resResp, &pb.SubjectsResponse{
			Token:      v.Token,
			SubjectsId: v.SubjectID,
		})
	}

	return &pb.LookupSubjectsResponse{
		SubjectsResponse: resResp,
	}, nil
}

// ReadSchema 读schema
func (s *server) ReadSchema(ctx context.Context, req *pb.ReadSchemaRequest) (*pb.ReadSchemaResponse, error) {

	res, err := s.uc.ReadSchema(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ReadSchemaResponse{
		Schema: res,
	}, nil
}

// WriteSchema 写schema
func (s *server) WriteSchema(ctx context.Context, req *pb.WriteSchemaRequest) (*pb.WriteSchemaResponse, error) {

	if req.Schema == "" {
		return nil, status.Error(codes.Aborted, "failed to write schema")
	}
	if err := s.uc.WriteScheme(ctx, schema); err != nil {
		return nil, err
	}
	return &pb.WriteSchemaResponse{}, nil
}
