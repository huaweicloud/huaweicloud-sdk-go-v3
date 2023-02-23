package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataMaskPolicyItem struct {

	// 访问数组
	Accesses *[]PolicyItemAccess `json:"accesses,omitempty"`

	// 条件数组
	Conditions *[]PolicyItemCondition `json:"conditions,omitempty"`

	DataMaskInfo *PolicyItemDataMaskInfo `json:"data_mask_info,omitempty"`

	// 是否支持传递
	DelegateAdmin *bool `json:"delegate_admin,omitempty"`

	// 用户组
	Groups *[]string `json:"groups,omitempty"`

	// 角色
	Roles *[]string `json:"roles,omitempty"`

	// 用户
	Users *[]string `json:"users,omitempty"`
}

func (o DataMaskPolicyItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataMaskPolicyItem struct{}"
	}

	return strings.Join([]string{"DataMaskPolicyItem", string(data)}, " ")
}
