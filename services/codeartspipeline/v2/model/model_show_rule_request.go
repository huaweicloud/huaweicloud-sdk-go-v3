package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRuleRequest Request Object
type ShowRuleRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 规则ID
	RuleId string `json:"rule_id"`

	// 项目ID
	CloudProjectId *string `json:"cloud_project_id,omitempty"`
}

func (o ShowRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowRuleRequest", string(data)}, " ")
}
