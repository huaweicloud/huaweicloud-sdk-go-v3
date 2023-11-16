package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOpenSourceStrategyRequest Request Object
type ShowOpenSourceStrategyRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 规则集ID
	RuleSetId string `json:"rule_set_id"`
}

func (o ShowOpenSourceStrategyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOpenSourceStrategyRequest struct{}"
	}

	return strings.Join([]string{"ShowOpenSourceStrategyRequest", string(data)}, " ")
}
