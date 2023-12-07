package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSecurityPermissionSetPermissionsRequest Request Object
type BatchDeleteSecurityPermissionSetPermissionsRequest struct {

	// 权限集id
	PermissionSetId string `json:"permission_set_id"`

	// workspace 信息
	Workspace string `json:"workspace"`

	Body *BatchDeleteBaseDto `json:"body,omitempty"`
}

func (o BatchDeleteSecurityPermissionSetPermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSecurityPermissionSetPermissionsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteSecurityPermissionSetPermissionsRequest", string(data)}, " ")
}
