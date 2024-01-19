package service

import (
	"context"

	pb "login_server/api/login_server/v1"
)

type LoginServerService struct {
	pb.UnimplementedLoginServerServer
}

func NewLoginServerService() *LoginServerService {
	return &LoginServerService{}
}

func (s *LoginServerService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
