package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MarketOptions 计费信息
type MarketOptions struct {
	ChargeMode *ChargeMode `json:"charge_mode,omitempty"`

	PrepaidOptions *PrepaidOptions `json:"prepaid_options,omitempty"`

	// 销售策略列表
	Strategies *[]Strategy `json:"strategies,omitempty"`
}

func (o MarketOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MarketOptions struct{}"
	}

	return strings.Join([]string{"MarketOptions", string(data)}, " ")
}
