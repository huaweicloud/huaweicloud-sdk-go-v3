package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RowFilterPolicyItem struct {

	// 权限
	Accesses *[]PolicyItemAccess `json:"accesses,omitempty"`

	// 条件
	Conditions *[]PolicyItemCondition `json:"conditions,omitempty"`

	// 是否传递
	DelegateAdmin *bool `json:"delegate_admin,omitempty"`

	// 组
	Groups *[]string `json:"groups,omitempty"`

	// 角色
	Roles *[]string `json:"roles,omitempty"`

	RowFilterInfo *PolicyItemRowFilterInfo `json:"row_filter_info,omitempty"`

	// 用户
	Users *[]string `json:"users,omitempty"`
}

func (o RowFilterPolicyItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RowFilterPolicyItem struct{}"
	}

	return strings.Join([]string{"RowFilterPolicyItem", string(data)}, " ")
}
