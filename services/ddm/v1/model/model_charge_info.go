package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChargeInfo struct {

	// 付费模式。
	ChargeMode *string `json:"charge_mode,omitempty"`

	// 使用周期。
	PeriodNum *int32 `json:"period_num,omitempty"`

	// 周期模式。
	PeriodType *string `json:"period_type,omitempty"`

	// 是否自动支付。
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`

	// 是否自动续费。
	IsAutoRenew *bool `json:"is_auto_renew,omitempty"`
}

func (o ChargeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChargeInfo struct{}"
	}

	return strings.Join([]string{"ChargeInfo", string(data)}, " ")
}
