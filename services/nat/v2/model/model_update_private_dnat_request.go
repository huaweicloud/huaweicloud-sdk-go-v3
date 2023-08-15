package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePrivateDnatRequest Request Object
type UpdatePrivateDnatRequest struct {

	// DNAT规则的ID。
	DnatRuleId string `json:"dnat_rule_id"`

	Body *UpdatePrivateDnatRequestBody `json:"body,omitempty"`
}

func (o UpdatePrivateDnatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateDnatRequest struct{}"
	}

	return strings.Join([]string{"UpdatePrivateDnatRequest", string(data)}, " ")
}
