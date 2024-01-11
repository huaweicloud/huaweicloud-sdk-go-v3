package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityPermissionSetPermissionRequest Request Object
type UpdateSecurityPermissionSetPermissionRequest struct {

	// 权限集id
	PermissionSetId string `json:"permission_set_id"`

	// 权限id
	PermissionId string `json:"permission_id"`

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *PermissionSetPermissionUpdateDto `json:"body,omitempty"`
}

func (o UpdateSecurityPermissionSetPermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityPermissionSetPermissionRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecurityPermissionSetPermissionRequest", string(data)}, " ")
}
