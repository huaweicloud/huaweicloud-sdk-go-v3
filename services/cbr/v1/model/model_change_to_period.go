package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeToPeriod struct {

	// 付费模式，当前仅可选择：pre_paid
	ChargingMode *string `json:"charging_mode,omitempty"`

	// 创建类型，按年(year)或者按月(month)
	PeriodType string `json:"period_type"`

	// 创建类型的数量，按年或按月的个数
	PeriodNum int32 `json:"period_num"`

	// 到期后是否自动续期，默认不续期
	IsAutoRenew *bool `json:"is_auto_renew,omitempty"`

	// 是否自动付费，默认为不自动付费
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`

	// 跳转URL
	ConsoleUrl *string `json:"console_url,omitempty"`

	// 资源列表
	VaultIds []string `json:"vault_ids"`
}

func (o ChangeToPeriod) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeToPeriod struct{}"
	}

	return strings.Join([]string{"ChangeToPeriod", string(data)}, " ")
}
