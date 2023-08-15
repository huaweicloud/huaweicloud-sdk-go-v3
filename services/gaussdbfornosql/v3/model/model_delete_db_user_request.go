package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDbUserRequest Request Object
type DeleteDbUserRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *RedisDeleteDbUserRequest `json:"body,omitempty"`
}

func (o DeleteDbUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDbUserRequest struct{}"
	}

	return strings.Join([]string{"DeleteDbUserRequest", string(data)}, " ")
}
