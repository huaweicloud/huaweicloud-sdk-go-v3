package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleInstanceContent struct {

	// 分组名称
	GroupName string `json:"group_name"`

	// 分组类型
	Type string `json:"type"`

	// 继承后的子策略是否可以修改阈值
	CanModifyWhenInherit *bool `json:"can_modify_when_inherit,omitempty"`

	// 规则属性列表
	Properties []RuleInstanceProperty `json:"properties"`
}

func (o RuleInstanceContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleInstanceContent struct{}"
	}

	return strings.Join([]string{"RuleInstanceContent", string(data)}, " ")
}
