package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateBotMRuleStatusRequest Request Object
type BatchUpdateBotMRuleStatusRequest struct {

	// policyId
	PolicyId string `json:"policy_id"`

	Body *BatchUpdateBotMRuleStatusRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateBotMRuleStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateBotMRuleStatusRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateBotMRuleStatusRequest", string(data)}, " ")
}
