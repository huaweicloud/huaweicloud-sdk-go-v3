package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStrategyResponse Response Object
type UpdateStrategyResponse struct {

	// 状态
	Status *bool `json:"status,omitempty"`

	// 策略ID
	RuleSetId      *string `json:"rule_set_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateStrategyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStrategyResponse struct{}"
	}

	return strings.Join([]string{"UpdateStrategyResponse", string(data)}, " ")
}
