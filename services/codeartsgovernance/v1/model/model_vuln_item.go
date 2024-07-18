package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VulnItem struct {

	// 漏洞项描述
	Desc *string `json:"desc,omitempty"`

	// 风险等级
	RiskLevel *string `json:"risk_level,omitempty"`

	// 漏洞问题列表
	VulnRules *[]VulnRule `json:"vuln_rules,omitempty"`
}

func (o VulnItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulnItem struct{}"
	}

	return strings.Join([]string{"VulnItem", string(data)}, " ")
}
