package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimCardsFlowVo struct {

	// 套餐实例ID
	Id *int64 `json:"id,omitempty" xml:"id"`

	// 账户ID
	AccountId *string `json:"account_id,omitempty" xml:"account_id"`

	// sim卡ID
	SimCardId *int64 `json:"sim_card_id,omitempty" xml:"sim_card_id"`

	// 套餐ID
	PricePlanId *string `json:"price_plan_id,omitempty" xml:"price_plan_id"`

	// 套餐名称
	PricePlanName *string `json:"price_plan_name,omitempty" xml:"price_plan_name"`

	// ICCID
	Iccid *string `json:"iccid,omitempty" xml:"iccid"`

	// 总流量(MB),两位小数
	FlowTotal *float64 `json:"flow_total,omitempty" xml:"flow_total"`

	// 已使用流量(MB),两位小数
	FlowUsed *float64 `json:"flow_used,omitempty" xml:"flow_used"`

	// 剩余流量(MB),两位小数
	FlowLeft *float64 `json:"flow_left,omitempty" xml:"flow_left"`
}

func (o SimCardsFlowVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimCardsFlowVo struct{}"
	}

	return strings.Join([]string{"SimCardsFlowVo", string(data)}, " ")
}
