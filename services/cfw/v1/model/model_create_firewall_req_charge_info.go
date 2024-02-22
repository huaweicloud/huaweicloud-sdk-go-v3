package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFirewallReqChargeInfo 计费类型信息，支持包年/包月和按需，默认为按需。
type CreateFirewallReqChargeInfo struct {

	// 计费模式。  取值范围：  prePaid：预付费，即包年/包月。 postPaid：后付费，即按需付费。
	ChargeMode string `json:"charge_mode"`

	// 订购周期类型。  取值范围：  month：包月。 year：包年。 说明：“charge_mode”为“prePaid”时生效，且为必选值。
	PeriodType *string `json:"period_type,omitempty"`

	// “charge_mode”为“prePaid”时生效，且为必选值，指定订购的时间。  取值范围：  当“period_type”为“month”时，取值为1~9。 当“period_type”为“year”时，取值为1~3。
	PeriodNum *int32 `json:"period_num,omitempty"`

	// 创建包周期实例时可指定，表示是否自动续订，续订的周期和原周期相同，且续订时会自动支付。  true，为自动续订。 false，为不自动续订，默认该方式。
	IsAutoRenew bool `json:"is_auto_renew"`

	// 创建包周期时可指定，表示是否自动从客户的账户中支付，此字段不影响自动续订的支付方式。  true，为自动支付。（会自动选择折扣和优惠券进行优惠，然后自动从客户华为云账户中支付），自动支付失败后会生成订单成功(该订单应付金额是优惠后金额)、但订单状态为“待支付”，等待客户手动支付(手动支付时，客户还可以修改系统自动选择的折扣和优惠券) false，为手动支付，默认该方式。（需要客户手动去支付，客户可以选择折扣和优惠券）
	IsAutoPay bool `json:"is_auto_pay"`
}

func (o CreateFirewallReqChargeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFirewallReqChargeInfo struct{}"
	}

	return strings.Join([]string{"CreateFirewallReqChargeInfo", string(data)}, " ")
}
