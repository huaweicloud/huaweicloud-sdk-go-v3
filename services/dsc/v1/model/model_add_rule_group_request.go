package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddRuleGroupRequest Request Object
type AddRuleGroupRequest struct {
	Body *RuleGroupRequest `json:"body,omitempty"`
}

func (o AddRuleGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddRuleGroupRequest struct{}"
	}

	return strings.Join([]string{"AddRuleGroupRequest", string(data)}, " ")
}
