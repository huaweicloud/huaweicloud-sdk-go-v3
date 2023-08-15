package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyItem struct {

	// 访问控制数据
	Accesses *[]PolicyItemAccess `json:"accesses,omitempty"`

	// 条件
	Conditions *[]PolicyItemCondition `json:"conditions,omitempty"`

	// 是否传递
	DelegateAdmin *bool `json:"delegate_admin,omitempty"`

	// 组
	Groups *[]string `json:"groups,omitempty"`

	// 角色
	Roles *[]string `json:"roles,omitempty"`

	// 用户
	Users *[]string `json:"users,omitempty"`
}

func (o PolicyItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyItem struct{}"
	}

	return strings.Join([]string{"PolicyItem", string(data)}, " ")
}
