package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 付费方式信息，与请求参数相同。
type ChargeInfoResult struct {
	// 计费模式。 取值范围：   - prePaid：预付费，即包年/包月。   - postPaid：后付费，即按需付费。

	ChargeMode *string `json:"charge_mode,omitempty"`
	// 订购周期类型。 “charge_mode”为“prePaid”时生效，且为必选值。 取值范围：   - month：包月。   - year：包年。

	PeriodType *string `json:"period_type,omitempty"`
	// “charge_mode”为“prePaid”时生效，且为必选值，指定订购的时间。 取值范围：   - 当“period_type”为“month”时，取值为1~9。   - 当“period_type”为“year”时，取值为1~3

	PeriodNum *string `json:"period_num,omitempty"`
	// 创建包周期实例时可指定，表示是否自动续订，续订的周期和原周期相同，且续订时会自动支付。 取值范围：   - true，表示自动续订。   - false，表示不自动续订，默认为该方式。

	IsAutoRenew *string `json:"is_auto_renew,omitempty"`
	// 创建包周期实例时可指定，表示是否自动从账户中支付，该字段不影响自动续订的支付方式。 取值范围：   - true，表示自动从账户中支付   - false，表示手动从账户中支付，默认为该支付方式。

	IsAutoPay *string `json:"is_auto_pay,omitempty"`
}

func (o ChargeInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChargeInfoResult struct{}"
	}

	return strings.Join([]string{"ChargeInfoResult", string(data)}, " ")
}
