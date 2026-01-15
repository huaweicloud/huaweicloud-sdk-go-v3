package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PeriodOrderResp 详情返回包周期信息体。
type PeriodOrderResp struct {

	// 订单状态 INITIALIZATION：初始化，AWAITING_AUDIT：待审核，AWAITING_REFUND：待退款，PROCESSING：处理中，CANCELED：已取消，COMPLETED：已完成，AWAITING_PAID：待付款，AWAITING_APPROVE：待审批，UNORDERED：未下单，FROZEN_CAN_RENEW：冻结可续费，FROZEN_CAN_UNSUBSCRIBE：冻结可退订，FROZEN_POLICE：公安冻结，FROZEN_CAN_RENEW_FREE：冻结可续费 可结束，FROZEN_CAN_FREE：冻结可结束，LOCK_RESOURCE：限制资源锁定
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
