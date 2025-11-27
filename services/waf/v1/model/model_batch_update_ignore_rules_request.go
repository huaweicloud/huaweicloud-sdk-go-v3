package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateIgnoreRulesRequest Request Object
type BatchUpdateIgnoreRulesRequest struct {
	Body *BatchUpdateIgnoreRuleRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateIgnoreRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateIgnoreRulesRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateIgnoreRulesRequest", string(data)}, " ")
}
