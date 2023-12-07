package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSecurityDataClassificationRuleRequest Request Object
type BatchDeleteSecurityDataClassificationRuleRequest struct {

	// workspace 信息
	Workspace string `json:"workspace"`

	Body *BatchDeleteRulesBaseDto `json:"body,omitempty"`
}

func (o BatchDeleteSecurityDataClassificationRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSecurityDataClassificationRuleRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteSecurityDataClassificationRuleRequest", string(data)}, " ")
}
