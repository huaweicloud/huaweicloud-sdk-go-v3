package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateJobRolePermissionRequest Request Object
type BatchUpdateJobRolePermissionRequest struct {
	Body *RolePermissionsRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateJobRolePermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateJobRolePermissionRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateJobRolePermissionRequest", string(data)}, " ")
}
