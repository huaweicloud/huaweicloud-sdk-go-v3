package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOpenSourceStrategyRequest Request Object
type DeleteOpenSourceStrategyRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 策略ID
	RuleSetId string `json:"rule_set_id"`
}

func (o DeleteOpenSourceStrategyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOpenSourceStrategyRequest struct{}"
	}

	return strings.Join([]string{"DeleteOpenSourceStrategyRequest", string(data)}, " ")
}
