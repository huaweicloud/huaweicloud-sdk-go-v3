package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateRulesetRequest struct {
	Body *Ruleset `json:"body,omitempty" xml:"body"`
}

func (o CreateRulesetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRulesetRequest struct{}"
	}

	return strings.Join([]string{"CreateRulesetRequest", string(data)}, " ")
}
