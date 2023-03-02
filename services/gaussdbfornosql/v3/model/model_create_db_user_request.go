package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDbUserRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *RedisCreateDbUserRequest `json:"body,omitempty"`
}

func (o CreateDbUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDbUserRequest struct{}"
	}

	return strings.Join([]string{"CreateDbUserRequest", string(data)}, " ")
}
