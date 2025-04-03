package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublicCloudServiceOrder 云服务信息
type PublicCloudServiceOrder struct {

	// 是否自动支付：下单订购后，是否自动从客户的华为云账户中支付，而不需要客户手动去进行支付；  1：是（会自动选择折扣和优惠券进行优惠，然后自动从客户华为云账户中支付），自动支付失败后会生成订单成功(该订单应付金额是优惠后金额)、但订单状态为“待支付”，等待客户手动支付(手动支付时，客户还可以修改系统自动选择的折扣和优惠券)。 0：否（需要客户手动去支付，客户可以选择折扣和优惠券）。 默认值为“0”。
	IsAutoPay *int32 `json:"is_auto_pay,omitempty"`

	// 订购周期类型： 2：月； 3：年； 6：一次性（chargingMode=1 一次性计费场景使用）
	PeriodType int32 `json:"period_type"`

	// 订购周期数 取值大于0；小于等于0会报错
	PeriodNum int32 `json:"period_num"`

	// 是否自动续订，为空时表示不自动续订； 1：自动续订 0：不自动续订（默认）
	IsAutoRenew *int32 `json:"is_auto_renew,omitempty"`

	// 订购数量； 取值大于0
	SubscriptionNum int32 `json:"subscription_num"`

	// 用户购买云服务产品的资源规格，详见[资源类型](metastudio_02_0042.xml)。
	ResourceSpecCode string `json:"resource_spec_code"`
}

func (o PublicCloudServiceOrder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicCloudServiceOrder struct{}"
	}

	return strings.Join([]string{"PublicCloudServiceOrder", string(data)}, " ")
}
