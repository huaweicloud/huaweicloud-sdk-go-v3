package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PriceItem struct {

	// 商品Id
	OfferingId *string `json:"offering_id,omitempty"`

	// 币种，USD
	Currency *string `json:"currency,omitempty"`

	// 官网价
	OfficialPrice *string `json:"official_price,omitempty"`

	// 计费模式，PERIOD：包年/包月、ON_DEMAND：按需、ONE_TIME：一次性、ON_DEMAND_PKG：按需套餐包
	ChargingMode *string `json:"charging_mode,omitempty"`

	// 销售周期类型，0：天 2：月 3：年 4：小时
	PeriodType *int32 `json:"period_type,omitempty"`

	// 销售周期数列表
	PeriodNums *[]int32 `json:"period_nums,omitempty"`

	// 计费因子编码
	BillingUsageFactor *string `json:"billing_usage_factor,omitempty"`

	// 定价策略模式，SIMPLE：简单定价、DIVISION：阶梯定价、CONDITION：条件选择定价
	PolicyMode *string `json:"policy_mode,omitempty"`

	DivisionPattern *DivisionPattern `json:"division_pattern,omitempty"`

	// 条件选择定价结构，policy_mode=CONDITION 时有效
	ConditionNodes *[]ConditionSelectPattern `json:"condition_nodes,omitempty"`
}

func (o PriceItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PriceItem struct{}"
	}

	return strings.Join([]string{"PriceItem", string(data)}, " ")
}
