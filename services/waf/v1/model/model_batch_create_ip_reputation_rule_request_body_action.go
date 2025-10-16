package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateIpReputationRuleRequestBodyAction action
type BatchCreateIpReputationRuleRequestBodyAction struct {

	// category
	Category *string `json:"category,omitempty"`
}

func (o BatchCreateIpReputationRuleRequestBodyAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateIpReputationRuleRequestBodyAction struct{}"
	}

	return strings.Join([]string{"BatchCreateIpReputationRuleRequestBodyAction", string(data)}, " ")
}
