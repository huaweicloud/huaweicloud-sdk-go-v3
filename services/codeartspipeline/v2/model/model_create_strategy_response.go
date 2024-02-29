package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStrategyResponse Response Object
type CreateStrategyResponse struct {

	// 状态
	Status *bool `json:"status,omitempty"`

	// 策略ID
	RuleSetId      *string `json:"rule_set_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateStrategyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStrategyResponse struct{}"
	}

	return strings.Join([]string{"CreateStrategyResponse", string(data)}, " ")
}
