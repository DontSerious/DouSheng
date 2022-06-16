package pack

import (
	"time"

	"github.com/DontSerious/DouSheng/kitex_gen/user"
)

func BuildBaseResponse(statusCode int64, statusMsg string) *user.BaseResponse {
	return &user.BaseResponse{
		StatusCode:  statusCode,
		StatusMsg:   statusMsg,
		ServiceTime: time.Now().Unix(),
	}
}
