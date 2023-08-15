package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RedisCreateDbUserRequest struct {

	// 需要创建的账号列表
	Users *[]RedisUserForCreation `json:"users,omitempty"`
}

func (o RedisCreateDbUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedisCreateDbUserRequest struct{}"
	}

	return strings.Join([]string{"RedisCreateDbUserRequest", string(data)}, " ")
}
