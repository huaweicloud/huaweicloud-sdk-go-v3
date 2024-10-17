package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimPoolVo struct {

	// 流量池标识
	Id *int64 `json:"id,omitempty"`

	// 流量池名称
	PoolName *string `json:"pool_name,omitempty"`

	// 套餐名称
	PricePlanName *string `json:"price_plan_name,omitempty"`

	// 套餐标识
	PricePlanId *string `json:"price_plan_id,omitempty"`

	// 生效时间
	EffectiveTime *sdktime.SdkTime `json:"effective_time,omitempty"`

	// 失效时间
	ExpiredTime *sdktime.SdkTime `json:"expired_time,omitempty"`

	// 账期
	BillingCycle *string `json:"billing_cycle,omitempty"`

	// 可用流量(查询账期所在月份)，单位MB
	FlowTotal *float64 `json:"flow_total,omitempty"`

	// 已用流量(查询账期所在月份), 单位MB
	FlowUsed *float64 `json:"flow_used,omitempty"`

	// 剩余流量(查询账期所在月份), 单位MB
	FlowLeft *float64 `json:"flow_left,omitempty"`

	// 流量池成员数量
	Quantity *int32 `json:"quantity,omitempty"`

	// 更新时间
	ModifyTime *sdktime.SdkTime `json:"modify_time,omitempty"`

	// 批次号
	OrderId *int64 `json:"order_id,omitempty"`

	// 已激活成员数量
	ActivatedSimQuantity *int32 `json:"activated_sim_quantity,omitempty"`

	// 未激活成员数量
	InactiveSimQuantity *int32 `json:"inactive_sim_quantity,omitempty"`

	// 已拆机成员数量
	DisassembledSimQuantity *int32 `json:"disassembled_sim_quantity,omitempty"`

	// 组成流量池的批次号列表
	OrderIds *string `json:"order_ids,omitempty"`

	// 流量池状态，0：废弃，1：在用
	Status *int32 `json:"status,omitempty"`
}

func (o SimPoolVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimPoolVo struct{}"
	}

	return strings.Join([]string{"SimPoolVo", string(data)}, " ")
}
