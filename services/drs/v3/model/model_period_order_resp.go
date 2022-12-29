package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 详情返回包周期信息体。
type PeriodOrderResp struct {

	// 订单状态
	Status *string `json:"status,omitempty"`

	// 订单ID
	OrderId *string `json:"order_id,omitempty"`

	// 计费模式
	ChargingMode *int32 `json:"charging_mode,omitempty"`

	// 订购包周期类型
	PeriodType *int32 `json:"period_type,omitempty"`

	// 订购周期数
	PeriodNum *int32 `json:"period_num,omitempty"`

	// 是否自动续费
	IsAutoRenew *int32 `json:"is_auto_renew,omitempty"`

	// 资源生效时间（即资源创建时间）
	EffTime *string `json:"eff_time,omitempty"`

	// 到期时间
	ExpTime *string `json:"exp_time,omitempty"`
}

func (o PeriodOrderResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeriodOrderResp struct{}"
	}

	return strings.Join([]string{"PeriodOrderResp", string(data)}, " ")
}
