package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetricList struct {

	// metric 名称
	MetricName *string `json:"metric_name,omitempty"`

	// 计算结果
	Values *[]interface{} `json:"values,omitempty"`
}

func (o MetricList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricList struct{}"
	}

	return strings.Join([]string{"MetricList", string(data)}, " ")
}
