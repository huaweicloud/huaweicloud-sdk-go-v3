package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 计费模式参数
type CreateNet2CloudPhoneServerRequestBodyExtendParam struct {

	// 计费类型  0 表示包周期
	ChargingMode int32 `json:"charging_mode"`

	// 订购周期类型 2 表示月 3 表示年
	PeriodType int32 `json:"period_type"`

	// 订购周期数 当订购周期为月时，取值范围[1, 9]。 当订购周期为年时，取值范围[1,10]
	PeriodNum int32 `json:"period_num"`

	// 是否自动付款。默认不自动付款。 1 表示自动付款 0 表示不自动付款
	IsAutoPay *int32 `json:"is_auto_pay,omitempty"`
}

func (o CreateNet2CloudPhoneServerRequestBodyExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNet2CloudPhoneServerRequestBodyExtendParam struct{}"
	}

	return strings.Join([]string{"CreateNet2CloudPhoneServerRequestBodyExtendParam", string(data)}, " ")
}
