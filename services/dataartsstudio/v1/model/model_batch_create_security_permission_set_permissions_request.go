package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateSecurityPermissionSetPermissionsRequest Request Object
type BatchCreateSecurityPermissionSetPermissionsRequest struct {

	// 权限集id
	PermissionSetId string `json:"permission_set_id"`

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *PermissionSetPermissionCreateDto `json:"body,omitempty"`
}

func (o BatchCreateSecurityPermissionSetPermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSecurityPermissionSetPermissionsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateSecurityPermissionSetPermissionsRequest", string(data)}, " ")
}
