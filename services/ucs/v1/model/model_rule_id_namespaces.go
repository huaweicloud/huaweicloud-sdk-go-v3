package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleIdNamespaces struct {

	// 权限策略id
	RuleIDs *[]string `json:"ruleIDs,omitempty"`

	// 权限策略涉及到的命名空间
	Namespaces *[]string `json:"namespaces,omitempty"`
}

func (o RuleIdNamespaces) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleIdNamespaces struct{}"
	}

	return strings.Join([]string{"RuleIdNamespaces", string(data)}, " ")
}
