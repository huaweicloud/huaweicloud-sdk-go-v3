package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePrivateSnatRequest struct {

	// SNAT规则的ID。
	SnatRuleId string `json:"snat_rule_id"`

	Body *UpdatePrivateSnatOptionBody `json:"body,omitempty"`
}

func (o UpdatePrivateSnatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateSnatRequest struct{}"
	}

	return strings.Join([]string{"UpdatePrivateSnatRequest", string(data)}, " ")
}
