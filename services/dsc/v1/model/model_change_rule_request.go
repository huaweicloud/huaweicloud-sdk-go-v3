package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeRuleRequest Request Object
type ChangeRuleRequest struct {
	Body *RuleChangeRequest `json:"body,omitempty"`
}

func (o ChangeRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeRuleRequest struct{}"
	}

	return strings.Join([]string{"ChangeRuleRequest", string(data)}, " ")
}
