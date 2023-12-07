package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSecurityPermissionSetMembersRequest Request Object
type BatchDeleteSecurityPermissionSetMembersRequest struct {

	// 权限集id
	PermissionSetId string `json:"permission_set_id"`

	// workspace 信息
	Workspace string `json:"workspace"`

	Body *BatchDeleteBaseDto `json:"body,omitempty"`
}

func (o BatchDeleteSecurityPermissionSetMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSecurityPermissionSetMembersRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteSecurityPermissionSetMembersRequest", string(data)}, " ")
}
