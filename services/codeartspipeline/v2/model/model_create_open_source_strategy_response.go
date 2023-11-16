package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOpenSourceStrategyResponse Response Object
type CreateOpenSourceStrategyResponse struct {

	// 状态
	Status *bool `json:"status,omitempty"`

	// 规则集ID
	RuleSetId      *string `json:"rule_set_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateOpenSourceStrategyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOpenSourceStrategyResponse struct{}"
	}

	return strings.Join([]string{"CreateOpenSourceStrategyResponse", string(data)}, " ")
}
