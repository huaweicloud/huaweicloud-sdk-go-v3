package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type CreateL7RuleRequestBody struct {
	Rule *CreateRuleOption `json:"rule"`
}

func (o CreateL7RuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7RuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateL7RuleRequestBody", string(data)}, " ")
}
