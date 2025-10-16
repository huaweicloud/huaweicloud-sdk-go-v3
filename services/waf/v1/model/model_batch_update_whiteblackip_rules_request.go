package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateWhiteblackipRulesRequest Request Object
type BatchUpdateWhiteblackipRulesRequest struct {
	Body *BatchUpdateWhiteBlackIpRuleRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateWhiteblackipRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateWhiteblackipRulesRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateWhiteblackipRulesRequest", string(data)}, " ")
}
