package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceRoleRequest Request Object
type ShowInstanceRoleRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceRoleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceRoleRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceRoleRequest", string(data)}, " ")
}
