package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateRulesetRequest struct {
	Body *Ruleset `json:"body,omitempty"`
}

func (o CreateRulesetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRulesetRequest struct{}"
	}

	return strings.Join([]string{"CreateRulesetRequest", string(data)}, " ")
}
