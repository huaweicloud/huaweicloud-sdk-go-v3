package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityRule 规则详情
type SecurityRule struct {
	Severity *SecurityRuleSeverity `json:"severity,omitempty"`

	Cve *SecurityRuleCve `json:"cve,omitempty"`
}

func (o SecurityRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityRule struct{}"
	}

	return strings.Join([]string{"SecurityRule", string(data)}, " ")
}
