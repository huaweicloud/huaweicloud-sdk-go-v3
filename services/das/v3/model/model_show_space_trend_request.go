package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSpaceTrendRequest Request Object
type ShowSpaceTrendRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// 开始时间（Unix timestamp），单位：毫秒
	StartTime int64 `json:"start_time"`

	// 结束时间（Unix timestamp），单位：毫秒
	EndTime int64 `json:"end_time"`

	// 指标名称
	MetricName string `json:"metric_name"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`
}

func (o ShowSpaceTrendRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSpaceTrendRequest struct{}"
	}

	return strings.Join([]string{"ShowSpaceTrendRequest", string(data)}, " ")
}
