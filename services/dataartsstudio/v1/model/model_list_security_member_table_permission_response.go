package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityMemberTablePermissionResponse Response Object
type ListSecurityMemberTablePermissionResponse struct {

	// 权限总数
	Total *int32 `json:"total,omitempty"`

	// 成员权限列表（包含权限集的和权限审批）
	MemberPermissionList *[]MemberPermission `json:"member_permission_list,omitempty"`
	HttpStatusCode       int                 `json:"-"`
}

func (o ListSecurityMemberTablePermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityMemberTablePermissionResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityMemberTablePermissionResponse", string(data)}, " ")
}
