package service

import (
	"context"

	"github.com/DontSerious/DouSheng/kitex_gen/user"
	"github.com/DontSerious/DouSheng/cmd/user/dal/db"
	"github.com/DontSerious/DouSheng/pkg/errno"
)

type AuthenticationService struct {
	ctx context.Context
}

func NewAuthenticationService(ctx context.Context) *AuthenticationService {
	return &AuthenticationService{
		ctx: ctx,
	}
}

// Token鉴权,成功返回user_id,否则返回0
func (s *AuthenticationService) Authentication(req *user.AuthenticationRequest) (user_id int64, statusCode int64, err error) {
	userName := req.Username
	users, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return 0, errno.ServiceErrCode, err
	}
	if len(users) == 0 {
		return 0, errno.TokenInvalidErrCode, errno.TokenErr
	}
	u := users[0]
	return int64(u.ID), errno.SuccessCode, nil
}
