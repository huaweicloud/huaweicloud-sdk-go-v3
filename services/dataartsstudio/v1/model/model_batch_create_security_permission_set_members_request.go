package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateSecurityPermissionSetMembersRequest Request Object
type BatchCreateSecurityPermissionSetMembersRequest struct {

	// 权限集id
	PermissionSetId string `json:"permission_set_id"`

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *PermissionSetMemberBatchCreateDto `json:"body,omitempty"`
}

func (o BatchCreateSecurityPermissionSetMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSecurityPermissionSetMembersRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateSecurityPermissionSetMembersRequest", string(data)}, " ")
}
