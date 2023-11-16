package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectOpenSourceStrategyRequest Request Object
type ShowProjectOpenSourceStrategyRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 策略ID
	RuleSetId string `json:"rule_set_id"`
}

func (o ShowProjectOpenSourceStrategyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectOpenSourceStrategyRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectOpenSourceStrategyRequest", string(data)}, " ")
}
