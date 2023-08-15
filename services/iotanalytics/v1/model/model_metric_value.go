package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetricValue struct {

	// 指标名称
	MetricName *string `json:"metric_name,omitempty"`

	// 计算结果，示例：[1,2]
	Values *[]interface{} `json:"values,omitempty"`
}

func (o MetricValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricValue struct{}"
	}

	return strings.Join([]string{"MetricValue", string(data)}, " ")
}
