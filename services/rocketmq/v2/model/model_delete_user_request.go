package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteUserRequest Request Object
type DeleteUserRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 用户名。
	UserName string `json:"user_name"`
}

func (o DeleteUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteUserRequest struct{}"
	}

	return strings.Join([]string{"DeleteUserRequest", string(data)}, " ")
}
