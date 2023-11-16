package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStrategyRequest Request Object
type DeleteStrategyRequest struct {

	// 规则集ID
	RuleSetId string `json:"rule_set_id"`

	// 租户ID
	DomainId string `json:"domain_id"`
}

func (o DeleteStrategyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStrategyRequest struct{}"
	}

	return strings.Join([]string{"DeleteStrategyRequest", string(data)}, " ")
}
