package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNatGatewaySnatRuleRequestOption 更新SNAT规则的请求体。
type UpdateNatGatewaySnatRuleRequestOption struct {
	SnatRule *UpdateNatGatewaySnatRuleOption `json:"snat_rule"`
}

func (o UpdateNatGatewaySnatRuleRequestOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatGatewaySnatRuleRequestOption struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewaySnatRuleRequestOption", string(data)}, " ")
}
