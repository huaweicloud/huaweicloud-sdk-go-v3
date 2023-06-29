package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRoleNamesRequest Request Object
type ListRoleNamesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ListRoleNamesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRoleNamesRequest struct{}"
	}

	return strings.Join([]string{"ListRoleNamesRequest", string(data)}, " ")
}
