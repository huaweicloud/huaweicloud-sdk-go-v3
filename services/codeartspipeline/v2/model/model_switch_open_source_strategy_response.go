package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchOpenSourceStrategyResponse Response Object
type SwitchOpenSourceStrategyResponse struct {

	// 状态
	Status *bool `json:"status,omitempty"`

	// 规则模版实例ID
	RuleTemplateInstanceId *string `json:"rule_template_instance_id,omitempty"`
	HttpStatusCode         int     `json:"-"`
}

func (o SwitchOpenSourceStrategyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchOpenSourceStrategyResponse struct{}"
	}

	return strings.Join([]string{"SwitchOpenSourceStrategyResponse", string(data)}, " ")
}
