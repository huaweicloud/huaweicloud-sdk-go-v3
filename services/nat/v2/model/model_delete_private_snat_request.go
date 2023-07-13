package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePrivateSnatRequest Request Object
type DeletePrivateSnatRequest struct {

	// SNAT规则的ID。
	SnatRuleId string `json:"snat_rule_id"`
}

func (o DeletePrivateSnatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateSnatRequest struct{}"
	}

	return strings.Join([]string{"DeletePrivateSnatRequest", string(data)}, " ")
}
