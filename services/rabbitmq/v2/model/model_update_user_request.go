package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserRequest Request Object
type UpdateUserRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 用户名。
	UserName string `json:"user_name"`

	Body *AmqpUser `json:"body,omitempty"`
}

func (o UpdateUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserRequest struct{}"
	}

	return strings.Join([]string{"UpdateUserRequest", string(data)}, " ")
}
