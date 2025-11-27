package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleNamespace struct {

	// 权限策略列表
	Rules *[]RuleInfo `json:"rules,omitempty"`

	// 命名空间列表
	Namespaces *[]string `json:"namespaces,omitempty"`
}

func (o RuleNamespace) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleNamespace struct{}"
	}

	return strings.Join([]string{"RuleNamespace", string(data)}, " ")
}
