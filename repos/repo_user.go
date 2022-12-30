package repos

import (
	"context"

	"dot.go/entities"
)

type ParamUser struct {
	ID    int32
	UA    string
	Token string
}

type RepoUser interface {
	UserLogin(ctx context.Context, user entities.User) (int32, error)
	UserDetail(ctx context.Context, id_user int32) (entities.User, error)
	UserGenerateToken(ctx context.Context, param ParamUser) (string, error)
	UserValidateToken(ctx context.Context, param ParamUser) (bool, error)
	UserUpdate(ctx context.Context, user entities.User) error
}
