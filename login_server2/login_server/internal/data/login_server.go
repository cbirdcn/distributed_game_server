package data

import (
	"context"

	"login_server/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type loginServerRepo struct {
	data *Data
	log  *log.Helper
}

// NewLoginServerRepo .
func NewLoginServerRepo(data *Data, logger log.Logger) biz.LoginServerRepo {
	return &loginServerRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *loginServerRepo) Save(ctx context.Context, g *biz.LoginServer) (*biz.LoginServer, error) {
	return g, nil
}

func (r *loginServerRepo) Update(ctx context.Context, g *biz.LoginServer) (*biz.LoginServer, error) {
	return g, nil
}

func (r *loginServerRepo) FindByID(context.Context, int64) (*biz.LoginServer, error) {
	return nil, nil
}

func (r *loginServerRepo) ListByHello(context.Context, string) ([]*biz.LoginServer, error) {
	return nil, nil
}

func (r *loginServerRepo) ListAll(context.Context) ([]*biz.LoginServer, error) {
	return nil, nil
}
