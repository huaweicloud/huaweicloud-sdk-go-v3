package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimCardsFlowVo struct {

	// 套餐实例ID
	Id *int64 `json:"id,omitempty"`

	// 账户ID
	AccountId *string `json:"account_id,omitempty"`

	// sim卡ID
	SimCardId *int64 `json:"sim_card_id,omitempty"`

	// 套餐ID
	PricePlanId *string `json:"price_plan_id,omitempty"`

	// 套餐名称
	PricePlanName *string `json:"price_plan_name,omitempty"`

	// ICCID
	Iccid *string `json:"iccid,omitempty"`

	// 总流量(MB),两位小数
	FlowTotal *float64 `json:"flow_total,omitempty"`

	// 已使用流量(MB),两位小数
	FlowUsed *float64 `json:"flow_used,omitempty"`

	// 剩余流量(MB),两位小数
	FlowLeft *float64 `json:"flow_left,omitempty"`
}

func (o SimCardsFlowVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimCardsFlowVo struct{}"
	}

	return strings.Join([]string{"SimCardsFlowVo", string(data)}, " ")
}
