package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IndicatorParam 自定义指标参数
type IndicatorParam struct {
	CustomizeParameter *CustomizeParameter `json:"customize_parameter,omitempty"`

	CustomizeFormula *CustomizeFormula `json:"customize_formula,omitempty"`
}

func (o IndicatorParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndicatorParam struct{}"
	}

	return strings.Join([]string{"IndicatorParam", string(data)}, " ")
}
