package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectStrategyRequest Request Object
type ShowProjectStrategyRequest struct {

	// 策略ID
	RuleSetId string `json:"rule_set_id"`

	// 项目ID
	ProjectId string `json:"project_id"`
}

func (o ShowProjectStrategyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectStrategyRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectStrategyRequest", string(data)}, " ")
}
