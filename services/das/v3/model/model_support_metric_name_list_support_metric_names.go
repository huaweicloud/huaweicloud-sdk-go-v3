package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SupportMetricNameListSupportMetricNames struct {

	// 数据库类型
	DatastoreType *[]string `json:"datastore_type,omitempty"`

	// 指标名称
	MetricName *string `json:"metric_name,omitempty"`

	// 单位
	Unit *string `json:"unit,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o SupportMetricNameListSupportMetricNames) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupportMetricNameListSupportMetricNames struct{}"
	}

	return strings.Join([]string{"SupportMetricNameListSupportMetricNames", string(data)}, " ")
}
