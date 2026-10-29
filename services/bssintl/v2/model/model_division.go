package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type Division struct {

	// 阶梯起始值
	BeginValue *decimal.Decimal `json:"begin_value,omitempty"`

	// 阶梯结束值
	EndValue *decimal.Decimal `json:"end_value,omitempty"`

	// 阶梯起始单位
	BeginUnit *int32 `json:"begin_unit,omitempty"`

	// 阶梯结束单位
	EndUnit *int32 `json:"end_unit,omitempty"`

	// 阶梯官网价
	OfficialPrice *string `json:"official_price,omitempty"`
}

func (o Division) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Division struct{}"
	}

	return strings.Join([]string{"Division", string(data)}, " ")
}
