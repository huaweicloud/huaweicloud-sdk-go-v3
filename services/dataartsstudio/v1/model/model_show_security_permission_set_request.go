package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityPermissionSetRequest Request Object
type ShowSecurityPermissionSetRequest struct {

	// 权限集id
	PermissionSetId string `json:"permission_set_id"`

	// workspace 信息
	Workspace string `json:"workspace"`
}

func (o ShowSecurityPermissionSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityPermissionSetRequest struct{}"
	}

	return strings.Join([]string{"ShowSecurityPermissionSetRequest", string(data)}, " ")
}
