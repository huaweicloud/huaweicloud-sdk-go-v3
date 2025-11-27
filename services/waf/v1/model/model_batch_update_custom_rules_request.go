package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateCustomRulesRequest Request Object
type BatchUpdateCustomRulesRequest struct {
	Body *BatchUpdateCustomRuleRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateCustomRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateCustomRulesRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateCustomRulesRequest", string(data)}, " ")
}
