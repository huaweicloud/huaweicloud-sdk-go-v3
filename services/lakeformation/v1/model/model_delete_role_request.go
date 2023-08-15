package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRoleRequest Request Object
type DeleteRoleRequest struct {

	// 实例Id
	InstanceId string `json:"instance_id"`

	// 角色名称
	RoleName string `json:"role_name"`
}

func (o DeleteRoleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRoleRequest struct{}"
	}

	return strings.Join([]string{"DeleteRoleRequest", string(data)}, " ")
}
