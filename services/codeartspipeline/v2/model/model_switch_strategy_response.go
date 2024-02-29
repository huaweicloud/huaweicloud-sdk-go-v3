package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchStrategyResponse Response Object
type SwitchStrategyResponse struct {

	// 状态
	Status *bool `json:"status,omitempty"`

	// 策略ID
	RuleSetId      *string `json:"rule_set_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchStrategyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchStrategyResponse struct{}"
	}

	return strings.Join([]string{"SwitchStrategyResponse", string(data)}, " ")
}
