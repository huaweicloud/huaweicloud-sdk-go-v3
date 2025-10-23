package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListComplianceRuleRequest Request Object
type ListComplianceRuleRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 是否系统内置
	EmbeddedFlag bool `json:"embedded_flag"`
}

func (o ListComplianceRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComplianceRuleRequest struct{}"
	}

	return strings.Join([]string{"ListComplianceRuleRequest", string(data)}, " ")
}
