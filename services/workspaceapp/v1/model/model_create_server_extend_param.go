package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateServerExtendParam struct {

	// 计费模式，取值范围： - prePaid-预付费，即包年包月； - postPaid-后付费，即按需付费；
	ChargingMode *string `json:"charging_mode,omitempty"`

	// 周期类型 2：包月；3：包年* chargingMode为prePaid时生效且为必选值;
	PeriodType *int32 `json:"period_type,omitempty"`

	// 订购周期数，chargingMode为prePaid时生效且为必选值；periodNum为正整数。取值范围： > - periodType=2（周期类型为月）时，取值为[1，9]； > - periodType=3（周期类型为年）时，取值为[1，3]；
	PeriodNum *int32 `json:"period_num,omitempty"`

	// 是否是自动续订，默认不填为false >- false 不自动续订 >- true 自动续订
	IsAutoRenew *bool `json:"is_auto_renew,omitempty"`

	// 下单订购后，是否自动从客户的账户中支付，而不需要客户手动去进行支付。chargingMode为prePaid时生效，不传该字段时默认为客户手动支付。 > - true ：是（自动支付） > - false：否（需要客户手动支付）
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o CreateServerExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServerExtendParam struct{}"
	}

	return strings.Join([]string{"CreateServerExtendParam", string(data)}, " ")
}
