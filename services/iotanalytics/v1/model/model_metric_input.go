package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricInput 定义指标计算查询的输入资产属性
type MetricInput struct {

	// 指标计算表达式的入参名称
	Name string `json:"name"`

	// 入参所对应的资产属性名称
	PropertyName string `json:"property_name"`
}

func (o MetricInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricInput struct{}"
	}

	return strings.Join([]string{"MetricInput", string(data)}, " ")
}
