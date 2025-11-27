package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRuleRequestBody struct {
	Spec *RuleSpec `json:"spec,omitempty"`
}

func (o UpdateRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateRuleRequestBody", string(data)}, " ")
}
