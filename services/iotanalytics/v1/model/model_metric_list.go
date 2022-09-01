package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetricList struct {

	// metric 名称
	MetricName *string `json:"metric_name,omitempty" xml:"metric_name"`

	// 计算结果
	Values *[]interface{} `json:"values,omitempty" xml:"values"`
}

func (o MetricList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricList struct{}"
	}

	return strings.Join([]string{"MetricList", string(data)}, " ")
}
