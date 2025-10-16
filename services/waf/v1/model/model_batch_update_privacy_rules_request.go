package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdatePrivacyRulesRequest Request Object
type BatchUpdatePrivacyRulesRequest struct {
	Body *BatchUpdatePrivacyRuleRequestBody `json:"body,omitempty"`
}

func (o BatchUpdatePrivacyRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePrivacyRulesRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdatePrivacyRulesRequest", string(data)}, " ")
}
