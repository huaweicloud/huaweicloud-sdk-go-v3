package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddRuleRequest Request Object
type AddRuleRequest struct {
	Body *RuleRequest `json:"body,omitempty"`
}

func (o AddRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddRuleRequest struct{}"
	}

	return strings.Join([]string{"AddRuleRequest", string(data)}, " ")
}
