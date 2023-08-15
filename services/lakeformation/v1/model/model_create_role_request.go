package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRoleRequest Request Object
type CreateRoleRequest struct {

	// 实例Id
	InstanceId string `json:"instance_id"`

	Body *RoleInput `json:"body,omitempty"`
}

func (o CreateRoleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRoleRequest struct{}"
	}

	return strings.Join([]string{"CreateRoleRequest", string(data)}, " ")
}
