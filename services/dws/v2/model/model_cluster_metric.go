package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterMetric struct {

	// 指标名称。
	Scope *string `json:"scope,omitempty"`

	// 指标表相关字段信息。
	Fields *[]SimpleFieldDto `json:"fields,omitempty"`

	// 作用域。
	MetricName *string `json:"metric_name,omitempty"`

	// 采集速率。
	CollectRate *int32 `json:"collect_rate,omitempty"`

	// 采集时间范围。
	CollectRange *[]string `json:"collect_range,omitempty"`

	// 创建时间。
	CreateTime *string `json:"create_time,omitempty"`
}

func (o ClusterMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterMetric struct{}"
	}

	return strings.Join([]string{"ClusterMetric", string(data)}, " ")
}
