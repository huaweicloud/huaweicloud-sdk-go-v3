package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivacyComplianceSubtype struct {

	// 隐私合规子类型描述
	Desc *string `json:"desc,omitempty"`

	// 检测项列表
	CheckerRules *[]VulCheckerRule `json:"checker_rules,omitempty"`
}

func (o PrivacyComplianceSubtype) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivacyComplianceSubtype struct{}"
	}

	return strings.Join([]string{"PrivacyComplianceSubtype", string(data)}, " ")
}
