package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimPoolVo struct {

	// 流量池标识
	Id *int64 `json:"id,omitempty" xml:"id"`

	// 流量池名称
	PoolName *string `json:"pool_name,omitempty" xml:"pool_name"`

	// 套餐名称
	PricePlanName *string `json:"price_plan_name,omitempty" xml:"price_plan_name"`

	// 套餐标识
	PricePlanId *string `json:"price_plan_id,omitempty" xml:"price_plan_id"`

	// 生效时间
	EffectiveTime *sdktime.SdkTime `json:"effective_time,omitempty" xml:"effective_time"`

	// 失效时间
	ExpiredTime *sdktime.SdkTime `json:"expired_time,omitempty" xml:"expired_time"`

	// 账期
	BillingCycle *string `json:"billing_cycle,omitempty" xml:"billing_cycle"`

	// 可用流量(查询账期所在月份)，单位MB
	FlowTotal *float64 `json:"flow_total,omitempty" xml:"flow_total"`

	// 已用流量(查询账期所在月份), 单位MB
	FlowUsed *float64 `json:"flow_used,omitempty" xml:"flow_used"`

	// 剩余流量(查询账期所在月份), 单位MB
	FlowLeft *float64 `json:"flow_left,omitempty" xml:"flow_left"`

	// 流量池成员数量
	Quantity *int32 `json:"quantity,omitempty" xml:"quantity"`

	// 更新时间
	ModifyTime *sdktime.SdkTime `json:"modify_time,omitempty" xml:"modify_time"`

	// 已激活成员数量
	ActivatedSimQuantity *int32 `json:"activated_sim_quantity,omitempty" xml:"activated_sim_quantity"`

	// 未激活成员数量
	InactiveSimQuantity *int32 `json:"inactive_sim_quantity,omitempty" xml:"inactive_sim_quantity"`

	// 已拆机成员数量
	DisassembledSimQuantity *int32 `json:"disassembled_sim_quantity,omitempty" xml:"disassembled_sim_quantity"`

	// 组成流量池的批次号列表
	OrderIds *string `json:"order_ids,omitempty" xml:"order_ids"`
}

func (o SimPoolVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimPoolVo struct{}"
	}

	return strings.Join([]string{"SimPoolVo", string(data)}, " ")
}
