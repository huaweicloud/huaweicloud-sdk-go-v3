package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RedisDeleteDbUserRequest 删除账号请求
type RedisDeleteDbUserRequest struct {

	// 需要删除的数据库账号名称列表。
	Names []string `json:"names"`
}

func (o RedisDeleteDbUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedisDeleteDbUserRequest struct{}"
	}

	return strings.Join([]string{"RedisDeleteDbUserRequest", string(data)}, " ")
}
