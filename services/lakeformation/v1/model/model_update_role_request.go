package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRoleRequest Request Object
type UpdateRoleRequest struct {

	// 实例Id
	InstanceId string `json:"instance_id"`

	// 角色名称
	RoleName string `json:"role_name"`

	Body *AlterRoleInput `json:"body,omitempty"`
}

func (o UpdateRoleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRoleRequest struct{}"
	}

	return strings.Join([]string{"UpdateRoleRequest", string(data)}, " ")
}
