package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStrategyRequest Request Object
type ShowStrategyRequest struct {

	// 策略ID
	RuleSetId string `json:"rule_set_id"`

	// 租户ID
	DomainId string `json:"domain_id"`

	// 项目ID
	CloudProjectId *string `json:"cloud_project_id,omitempty"`
}

func (o ShowStrategyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStrategyRequest struct{}"
	}

	return strings.Join([]string{"ShowStrategyRequest", string(data)}, " ")
}
