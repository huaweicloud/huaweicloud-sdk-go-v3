package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChargeInfoVo struct {

	// 计费模式，取值：  - on_demand: 按需  - period: 包周期
	ChargeMode *string `json:"charge_mode,omitempty"`

	// 订购周期类型，“charge_mode”为“period”时生效，且为必选值。取值：  - month: 包月  - year: 包年
	PeriodType *string `json:"period_type,omitempty"`

	// 订购时长，“charge_mode”为“period”时生效，且为必选值，指定订购的时间。 当“period_type”为“month”时，取值为1~9。 当“period_type”为“year”时，取值为1~3和5。
	PeriodNum *int32 `json:"period_num,omitempty"`

	// 是否自动续订，创建包周期实例时可指定，表示是否自动续订，续订的周期和原周期相同，且续订时会自动支付。取值：  - true: 自动续订  - false: 不自动续订
	IsAutoRenew *bool `json:"is_auto_renew,omitempty"`

	// 是否自动支付，创建包周期时可指定，表示是否自动从客户的账户中支付，此字段不影响自动续订的支付方式。  - true: 自动支付  - false: 手动支付
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o ChargeInfoVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChargeInfoVo struct{}"
	}

	return strings.Join([]string{"ChargeInfoVo", string(data)}, " ")
}
