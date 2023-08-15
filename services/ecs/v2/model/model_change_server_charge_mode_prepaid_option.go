package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeServerChargeModePrepaidOption struct {

	// 是否连同支持的按需数据盘一起转为包周期。 当参数为true，包括按需非共享的数据盘，不包括共享云硬盘，DSS和DESS硬盘 默认值为false
	IncludeDataDisks *bool `json:"include_data_disks,omitempty"`

	// 是否连同弹性公网IP一起转为包周期 只有“独享”、“按带宽计费”的弹性公网IP才可以转换为包周期计费模式 默认值为false
	IncludePublicips *bool `json:"include_publicips,omitempty"`

	// 订购周期类型，取值范围： month-月 year-年
	PeriodType string `json:"period_type"`

	// 订购周期的周期数。 取值范围： period_type=month时，取值范围为[1,9]。 period_type=year时，取值范围为[1,3]。
	PeriodNum string `json:"period_num"`

	// 是否自动支付。 true：自动支付，需要确保账户余额充足，如果余额不足则会生成异常订单，只能作废此订单 false：只生成订单不扣费 默认值为false
	AutoPay *bool `json:"auto_pay,omitempty"`

	// 是否自动续费。默认值：false
	AutoRenew *bool `json:"auto_renew,omitempty"`
}

func (o ChangeServerChargeModePrepaidOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeServerChargeModePrepaidOption struct{}"
	}

	return strings.Join([]string{"ChangeServerChargeModePrepaidOption", string(data)}, " ")
}
