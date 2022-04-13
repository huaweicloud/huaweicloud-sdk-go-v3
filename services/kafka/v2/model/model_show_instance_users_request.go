package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowInstanceUsersRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceUsersRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceUsersRequest", string(data)}, " ")
}
