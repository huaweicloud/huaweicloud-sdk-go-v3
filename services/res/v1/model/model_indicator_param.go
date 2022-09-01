package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 自定义指标参数
type IndicatorParam struct {
	CustomizeParameter *CustomizeParameter `json:"customize_parameter,omitempty" xml:"customize_parameter"`

	CustomizeFormula *CustomizeFormula `json:"customize_formula,omitempty" xml:"customize_formula"`
}

func (o IndicatorParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndicatorParam struct{}"
	}

	return strings.Join([]string{"IndicatorParam", string(data)}, " ")
}
