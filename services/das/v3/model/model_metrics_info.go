package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricsInfo 指标信息
type MetricsInfo struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 指标
	Series *[]float64 `json:"series,omitempty"`

	// 时间戳
	Timestamps *[]int64 `json:"timestamps,omitempty"`
}

func (o MetricsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricsInfo struct{}"
	}

	return strings.Join([]string{"MetricsInfo", string(data)}, " ")
}
