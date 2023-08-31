package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FlowUsageVo struct {

	// 账期
	BillingCycle *string `json:"billing_cycle,omitempty"`

	// 已用流量（MB）
	FlowUsed *float64 `json:"flow_used,omitempty"`
}

func (o FlowUsageVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlowUsageVo struct{}"
	}

	return strings.Join([]string{"FlowUsageVo", string(data)}, " ")
}
