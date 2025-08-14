package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetPermissionSetSummaryResponse Response Object
type GetPermissionSetSummaryResponse struct {

	// 已创建的权限集数量
	PermissionSets *int64 `json:"permission_sets,omitempty"`

	// 权限集配额
	PermissionSetsQuota *int64 `json:"permission_sets_quota,omitempty"`

	// 允许权限集可绑定的策略配额
	PermissionPolicyCountQuota *int64 `json:"permission_policy_count_quota,omitempty"`

	// 允许权限集可绑定的身份策略配额
	PermissionPolicyQuota *int64 `json:"permission_policy_quota,omitempty"`
	HttpStatusCode        int    `json:"-"`
}

func (o GetPermissionSetSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetPermissionSetSummaryResponse struct{}"
	}

	return strings.Join([]string{"GetPermissionSetSummaryResponse", string(data)}, " ")
}
