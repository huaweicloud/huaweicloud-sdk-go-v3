package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityMemberPermissionResponse Response Object
type ListSecurityMemberPermissionResponse struct {

	// 权限总数
	Total *int64 `json:"total,omitempty"`

	// 成员权限列表（包含权限集的和权限审批）
	Result         *[]AccountPermission `json:"result,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListSecurityMemberPermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityMemberPermissionResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityMemberPermissionResponse", string(data)}, " ")
}
