package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExceptionMetricsRequestBody Exception Analyze Query Metrics New请求体
type ListExceptionMetricsRequestBody struct {

	// 开始时间（Unix timestamp，毫秒）
	StartTime int64 `json:"start_time"`

	// 结束时间（Unix timestamp，毫秒）
	EndTime int64 `json:"end_time"`

	// 数据库引擎类型
	EngineType *string `json:"engine_type,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// CES指标名列表
	MetricNames []string `json:"metric_names"`

	// 聚合粒度
	Interval string `json:"interval"`

	// 聚合方式
	AggregationMode string `json:"aggregation_mode"`
}

func (o ListExceptionMetricsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExceptionMetricsRequestBody struct{}"
	}

	return strings.Join([]string{"ListExceptionMetricsRequestBody", string(data)}, " ")
}
