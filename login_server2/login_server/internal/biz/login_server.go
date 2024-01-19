package biz

import (
	"context"

	v1 "login_server/api/login_server/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// LoginServer is a LoginServer model.
type LoginServer struct {
	Hello string
}

// LoginServerRepo is a Greater repo.
type LoginServerRepo interface {
	Save(context.Context, *LoginServer) (*LoginServer, error)
	Update(context.Context, *LoginServer) (*LoginServer, error)
	FindByID(context.Context, int64) (*LoginServer, error)
	ListByHello(context.Context, string) ([]*LoginServer, error)
	ListAll(context.Context) ([]*LoginServer, error)
}

// LoginServerUsecase is a LoginServer usecase.
type LoginServerUsecase struct {
	repo LoginServerRepo
	log  *log.Helper
}

// NewLoginServerUsecase new a LoginServer usecase.
func NewLoginServerUsecase(repo LoginServerRepo, logger log.Logger) *LoginServerUsecase {
	return &LoginServerUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateLoginServer creates a LoginServer, and returns the new LoginServer.
func (uc *LoginServerUsecase) CreateLoginServer(ctx context.Context, g *LoginServer) (*LoginServer, error) {
	uc.log.WithContext(ctx).Infof("CreateLoginServer: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
