package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOpenSourceStrategyResponse Response Object
type UpdateOpenSourceStrategyResponse struct {

	// 状态
	Status *bool `json:"status,omitempty"`

	// 规则集ID
	RuleSetId      *string `json:"rule_set_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateOpenSourceStrategyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOpenSourceStrategyResponse struct{}"
	}

	return strings.Join([]string{"UpdateOpenSourceStrategyResponse", string(data)}, " ")
}
