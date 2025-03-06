package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCloudPhoneSingleServerRequestBodyOrderParam 订单结构体。
type CreateCloudPhoneSingleServerRequestBodyOrderParam struct {

	// 计费类型。 0：包周期
	ChargingMode int32 `json:"charging_mode"`

	// 订购周期类型。 2：月 3：年
	PeriodType int32 `json:"period_type"`

	// 订购周期数。 当订购周期为月时，取值范围[1, 9]。 当订购周期为年时，取值范围[1,10]。
	PeriodNum int32 `json:"period_num"`

	// 是否自动付款。默认不自动付款。 0：不自动付款 1：自动付款
	IsAutoPay *int32 `json:"is_auto_pay,omitempty"`

	// 是否自动续订。默认不自动续订。 0：不自动续订 1：自动续订
	IsAutoRenew *int32 `json:"is_auto_renew,omitempty"`
}

func (o CreateCloudPhoneSingleServerRequestBodyOrderParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudPhoneSingleServerRequestBodyOrderParam struct{}"
	}

	return strings.Join([]string{"CreateCloudPhoneSingleServerRequestBodyOrderParam", string(data)}, " ")
}
