package service

import (
	"context"

	v1 "login_server/api/helloworld/v1"
	"login_server/internal/biz"
)

// LoginServerService is a loginServer service.
type LoginServerService struct {
	v1.UnimplementedLoginServerServer

	uc *biz.LoginServerUsecase
}

// NewLoginServerService new a loginServer service.
func NewLoginServerService(uc *biz.LoginServerUsecase) *LoginServerService {
	return &LoginServerService{uc: uc}
}

// SayHello implements helloworld.LoginServerServer.
func (s *LoginServerService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateLoginServer(ctx, &biz.LoginServer{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
