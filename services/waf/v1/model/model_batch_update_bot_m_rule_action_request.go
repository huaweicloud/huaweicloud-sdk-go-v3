package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateBotMRuleActionRequest Request Object
type BatchUpdateBotMRuleActionRequest struct {

	// policyId
	PolicyId string `json:"policy_id"`

	Body *BatchUpdateBotMRuleActionRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateBotMRuleActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateBotMRuleActionRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateBotMRuleActionRequest", string(data)}, " ")
}
