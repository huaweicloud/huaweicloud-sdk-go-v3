package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetDbUserPasswordRequest Request Object
type ResetDbUserPasswordRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *RedisResetDbUserPasswordRequestBody `json:"body,omitempty"`
}

func (o ResetDbUserPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetDbUserPasswordRequest struct{}"
	}

	return strings.Join([]string{"ResetDbUserPasswordRequest", string(data)}, " ")
}
