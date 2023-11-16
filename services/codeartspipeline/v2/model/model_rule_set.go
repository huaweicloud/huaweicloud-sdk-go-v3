package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleSet struct {

	// 规则模版实例ID
	Id string `json:"id"`

	// 规则模版实例名称
	Name string `json:"name"`

	// 类型
	Type string `json:"type"`

	// 版本
	Version *string `json:"version,omitempty"`

	// 最近操作人
	Operator string `json:"operator"`

	// 最近操作时间
	OperateTime int64 `json:"operate_time"`

	// 是否生效
	IsValid bool `json:"is_valid"`

	// 租户级、项目级
	Level *string `json:"level,omitempty"`

	// 是否系统级
	IsPublic *bool `json:"is_public,omitempty"`
}

func (o RuleSet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleSet struct{}"
	}

	return strings.Join([]string{"RuleSet", string(data)}, " ")
}
