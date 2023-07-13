package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRoleRequest Request Object
type ShowRoleRequest struct {

	// 实例Id
	InstanceId string `json:"instance_id"`

	// 角色名称
	RoleName string `json:"role_name"`
}

func (o ShowRoleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRoleRequest struct{}"
	}

	return strings.Join([]string{"ShowRoleRequest", string(data)}, " ")
}
