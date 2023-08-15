package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChargeInfo struct {

	// 计费模式。 取值范围： - prePaid：预付费，即包年/包月。（创建集群接口现已支持预付费，创建集群并提交作业接口暂不支持预付费。） - postPaid：后付费，即按需计费。
	ChargeMode string `json:"charge_mode"`

	// 周期类型。取值范围： - month：包月。 - year： 包年。 - day：按需计费。
	PeriodType *string `json:"period_type,omitempty"`

	// 周期数，“charge_mode”为“prePaid”时生效，且为必选值，指定订购的时间。取值范围： - 当“period_type”为“month”时，取值为1~9。 - 当“period_type”为“year”时，取值为1~3。
	PeriodNum *int32 `json:"period_num,omitempty"`

	// 是否自动支付，包周期模式下使用，下单订购后，是否自动从客户的账户中支付，而不需要客户手动去进行支付，默认为手动支付。取值范围： - true：自动支付，会自动选择折扣和优惠券进行优惠，然后自动从客户账户中支付，自动支付失败后会生成订单成功、但订单状态为“待支付”，等待客户手动支付。 - false：手动支付，需要客户手动去支付，客户可以选择折扣和优惠券。
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o ChargeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChargeInfo struct{}"
	}

	return strings.Join([]string{"ChargeInfo", string(data)}, " ")
}
