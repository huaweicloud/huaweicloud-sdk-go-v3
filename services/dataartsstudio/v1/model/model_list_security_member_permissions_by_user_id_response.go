package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityMemberPermissionsByUserIdResponse Response Object
type ListSecurityMemberPermissionsByUserIdResponse struct {

	// 权限总数
	Total *int64 `json:"total,omitempty"`

	// 成员权限列表（包含权限集的和权限审批）
	Result         *[]AccountPermission `json:"result,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListSecurityMemberPermissionsByUserIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityMemberPermissionsByUserIdResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityMemberPermissionsByUserIdResponse", string(data)}, " ")
}
