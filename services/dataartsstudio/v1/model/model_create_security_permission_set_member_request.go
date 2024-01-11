package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecurityPermissionSetMemberRequest Request Object
type CreateSecurityPermissionSetMemberRequest struct {

	// 权限集id
	PermissionSetId string `json:"permission_set_id"`

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *PermissionSetMemberCreateDto `json:"body,omitempty"`
}

func (o CreateSecurityPermissionSetMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityPermissionSetMemberRequest struct{}"
	}

	return strings.Join([]string{"CreateSecurityPermissionSetMemberRequest", string(data)}, " ")
}
