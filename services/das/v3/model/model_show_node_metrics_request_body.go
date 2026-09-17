package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNodeMetricsRequestBody Query请求体
type ShowNodeMetricsRequestBody struct {

	// CES指标名列表
	MetricNames []string `json:"metric_names"`

	// 开始时间，Unix timestamp，单位：毫秒
	StartTime int64 `json:"start_time"`

	// 结束时间，Unix timestamp，单位：毫秒
	EndTime int64 `json:"end_time"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`
}

func (o ShowNodeMetricsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodeMetricsRequestBody struct{}"
	}

	return strings.Join([]string{"ShowNodeMetricsRequestBody", string(data)}, " ")
}
