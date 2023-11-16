package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOpenSourceStrategyRequest Request Object
type UpdateOpenSourceStrategyRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 规则集ID
	RuleSetId string `json:"rule_set_id"`

	Body *OpenSourceRuleSetCreateVo `json:"body,omitempty"`
}

func (o UpdateOpenSourceStrategyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOpenSourceStrategyRequest struct{}"
	}

	return strings.Join([]string{"UpdateOpenSourceStrategyRequest", string(data)}, " ")
}
