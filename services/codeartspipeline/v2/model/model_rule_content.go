package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleContent struct {

	// 分组名称
	GroupName string `json:"group_name"`

	// 继承后的子策略是否可以修改阈值
	CanModifyWhenInherit *bool `json:"can_modify_when_inherit,omitempty"`

	// 规则属性列表
	Properties []RuleProperty `json:"properties"`
}

func (o RuleContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleContent struct{}"
	}

	return strings.Join([]string{"RuleContent", string(data)}, " ")
}
