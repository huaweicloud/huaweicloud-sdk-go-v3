package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConditionSelectPattern struct {

	// 条件类型，SERVICE_HOUR_TIME：时间段 PERIOD_NUM：周期数 PERIOD_TYPE：周期类型，年、月、天、小时 PURCHASE_SCENE：订购场景，订购、续订、变更 NUMBER_OF_PEOPLE：线性大小
	ConditionType *string `json:"condition_type,omitempty"`

	// 条件表达式
	ConditionStatement *string `json:"condition_statement,omitempty"`

	// 定价策略类型，SIMPLE：简单定价 DIVISION：阶梯定价
	PolicyMode *string `json:"policy_mode,omitempty"`

	// 阶梯官网价，policy_mode取值为SIMPLE时为官网价，如果取值为DIVISION，则价格在division_pattern中。
	OfficialPrice *string `json:"official_price,omitempty"`

	DivisionPattern *DivisionPattern `json:"division_pattern,omitempty"`
}

func (o ConditionSelectPattern) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConditionSelectPattern struct{}"
	}

	return strings.Join([]string{"ConditionSelectPattern", string(data)}, " ")
}
