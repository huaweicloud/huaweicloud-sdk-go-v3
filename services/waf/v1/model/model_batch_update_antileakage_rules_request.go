package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateAntileakageRulesRequest Request Object
type BatchUpdateAntileakageRulesRequest struct {
	Body *BatchUpdateAntileakageRuleRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateAntileakageRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateAntileakageRulesRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateAntileakageRulesRequest", string(data)}, " ")
}
