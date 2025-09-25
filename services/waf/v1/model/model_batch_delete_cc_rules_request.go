package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteCcRulesRequest Request Object
type BatchDeleteCcRulesRequest struct {
	Body *PolicyRuleIdRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteCcRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCcRulesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteCcRulesRequest", string(data)}, " ")
}
