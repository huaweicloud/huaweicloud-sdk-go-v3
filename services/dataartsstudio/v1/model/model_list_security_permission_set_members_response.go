package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityPermissionSetMembersResponse Response Object
type ListSecurityPermissionSetMembersResponse struct {

	// 总条数
	Total *int32 `json:"total,omitempty"`

	// 成员列表
	PermissionSetMembers *[]PermissionSetMember `json:"permission_set_members,omitempty"`
	HttpStatusCode       int                    `json:"-"`
}

func (o ListSecurityPermissionSetMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityPermissionSetMembersResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityPermissionSetMembersResponse", string(data)}, " ")
}
