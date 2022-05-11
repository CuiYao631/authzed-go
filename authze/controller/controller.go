package controller

import (
	"authzed-go/authze/pb"
	"authzed-go/authze/usecase"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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

	return &pb.AddPermResponse{}, nil
}

// UpdatePerm 更新权限
func (s *server) UpdatePerm(ctx context.Context, req *pb.UpdatePermRequest) (*pb.UpdatePermResponse, error) {

	return &pb.UpdatePermResponse{}, nil
}

// DelPerm 删除权限
func (s *server) DelPerm(ctx context.Context, req *pb.DelPermRequest) (*pb.DelPermResponse, error) {

	return &pb.DelPermResponse{}, nil
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
	if err := s.uc.WriteScheme(ctx, req.Schema); err != nil {
		return nil, err
	}
	return &pb.WriteSchemaResponse{}, nil
}
