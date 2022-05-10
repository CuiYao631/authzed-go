package authze

import (
	"authzed-go/authze/controller"
	"authzed-go/authze/pb"
	"authzed-go/authze/usecase"
	"github.com/authzed/authzed-go/v1"
	"google.golang.org/grpc"
)

type authze struct {
	server controller.Server
}

func NewAuthzed(client *authzed.Client) *authze {
	uc := usecase.NewUsecase(client)
	s := controller.NewServer(uc)
	return &authze{s}
}
func (c *authze) InitServer(s *grpc.Server) {
	pb.RegisterAuthzedSVCServer(s, c.server)
}
