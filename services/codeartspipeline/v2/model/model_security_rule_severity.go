package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityRuleSeverity 漏洞级别
type SecurityRuleSeverity struct {

	// 是否启用
	Enable *bool `json:"enable,omitempty"`

	// 漏洞等级
	Values *[]string `json:"values,omitempty"`
}

func (o SecurityRuleSeverity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityRuleSeverity struct{}"
	}

	return strings.Join([]string{"SecurityRuleSeverity", string(data)}, " ")
}
