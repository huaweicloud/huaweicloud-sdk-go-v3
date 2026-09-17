package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricNamesSupportItem MetricNamesSupportItem对象
type MetricNamesSupportItem struct {

	// 数据库类型
	EngineTypes *[]string `json:"engine_types,omitempty"`

	// 指标名称
	MetricName *string `json:"metric_name,omitempty"`

	// 单位
	Unit *string `json:"unit,omitempty"`

	// 描述
	MetricNameDes *string `json:"metric_name_des,omitempty"`
}

func (o MetricNamesSupportItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricNamesSupportItem struct{}"
	}

	return strings.Join([]string{"MetricNamesSupportItem", string(data)}, " ")
}
