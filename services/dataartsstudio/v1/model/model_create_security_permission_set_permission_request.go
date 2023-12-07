package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecurityPermissionSetPermissionRequest Request Object
type CreateSecurityPermissionSetPermissionRequest struct {

	// 权限集id
	PermissionSetId string `json:"permission_set_id"`

	// workspace 信息
	Workspace string `json:"workspace"`

	Body *PermissionSetPermissionCreateDto `json:"body,omitempty"`
}

func (o CreateSecurityPermissionSetPermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityPermissionSetPermissionRequest struct{}"
	}

	return strings.Join([]string{"CreateSecurityPermissionSetPermissionRequest", string(data)}, " ")
}
