package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type UpdateL7RuleRequestBody struct {
	Rule *UpdateL7RuleOption `json:"rule"`
}

func (o UpdateL7RuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7RuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateL7RuleRequestBody", string(data)}, " ")
}
