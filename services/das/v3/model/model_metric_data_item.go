package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricDataItem 指标数据项
type MetricDataItem struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 指标值
	Series *[]float64 `json:"series,omitempty"`

	// 时间戳
	Timestamps *[]int64 `json:"timestamps,omitempty"`
}

func (o MetricDataItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricDataItem struct{}"
	}

	return strings.Join([]string{"MetricDataItem", string(data)}, " ")
}
