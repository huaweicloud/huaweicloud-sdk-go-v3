package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleSpec struct {

	// 权限策略关联的IAM用户信息
	IamUserIDs *[]string `json:"iamUserIDs,omitempty"`

	// 权限策略类型，只允许四种类型：readonly/develop/admin/custom
	Type *string `json:"type,omitempty"`

	// 权限策略内容
	Contents *[]Content `json:"contents,omitempty"`

	// 权限策略描述信息
	Description *string `json:"description,omitempty"`
}

func (o RuleSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleSpec struct{}"
	}

	return strings.Join([]string{"RuleSpec", string(data)}, " ")
}
