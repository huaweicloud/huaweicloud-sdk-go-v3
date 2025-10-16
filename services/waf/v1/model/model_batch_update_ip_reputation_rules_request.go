package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateIpReputationRulesRequest Request Object
type BatchUpdateIpReputationRulesRequest struct {
	Body *BatchUpdateIpReputationRuleRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateIpReputationRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateIpReputationRulesRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateIpReputationRulesRequest", string(data)}, " ")
}
