package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceOrderReq struct {

	// 产品编号
	ProductId *string `json:"product_id,omitempty"`

	// 计费模式。
	ChargingMode *int32 `json:"charging_mode,omitempty"`

	// 支付模式。
	PaymentMode *string `json:"payment_mode,omitempty"`

	// 订购周期类型： - 2：月 - 3：年
	PeriodType *int32 `json:"period_type,omitempty"`

	// 订购周期数
	PeriodNum *int32 `json:"period_num,omitempty"`

	// 是否支持自动续费
	IsAutoRenew *int32 `json:"is_auto_renew,omitempty"`

	// 促销产品编号
	PromotionId *string `json:"promotion_id,omitempty"`

	// 促销计划编号
	PromotionPlanId *string `json:"promotion_plan_id,omitempty"`

	// 促销信息
	PromotionInfo *string `json:"promotion_info,omitempty"`

	// 组合产品编号
	CompositeProductId *string `json:"composite_product_id,omitempty"`

	InstanceInfo *InstanceCreateReqV2 `json:"instance_info,omitempty"`
}

func (o InstanceOrderReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceOrderReq struct{}"
	}

	return strings.Join([]string{"InstanceOrderReq", string(data)}, " ")
}
