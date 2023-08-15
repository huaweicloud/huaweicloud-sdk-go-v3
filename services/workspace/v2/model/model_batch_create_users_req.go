package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateUsersReq struct {

	// 创建单个用户请求。
	Users []CreateUserRequest `json:"users"`
}

func (o BatchCreateUsersReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateUsersReq struct{}"
	}

	return strings.Join([]string{"BatchCreateUsersReq", string(data)}, " ")
}
