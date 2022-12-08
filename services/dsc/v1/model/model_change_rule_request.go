package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ChangeRuleRequest struct {
	Body *RuleRequest `json:"body,omitempty"`
}

func (o ChangeRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeRuleRequest struct{}"
	}

	return strings.Join([]string{"ChangeRuleRequest", string(data)}, " ")
}
