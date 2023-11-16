package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OpenSourceRuleContent 开源治理策略规则详情
type OpenSourceRuleContent struct {
	VersionSet *VersionSetProperty `json:"version_set,omitempty"`

	Security *SecurityProperty `json:"security,omitempty"`

	License *LicenseProperty `json:"license,omitempty"`
}

func (o OpenSourceRuleContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenSourceRuleContent struct{}"
	}

	return strings.Join([]string{"OpenSourceRuleContent", string(data)}, " ")
}
