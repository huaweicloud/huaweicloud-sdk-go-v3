package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TopRiskInfo struct {

	// 实例id。
	InstanceId string `json:"instance_id"`

	// 实例名称。
	InstanceName string `json:"instance_name"`

	// 节点ID。
	NodeId string `json:"node_id"`

	// 指标名称。
	MetricNames []string `json:"metric_names"`

	// 指标值,单位%。
	MetricValues []float64 `json:"metric_values"`
}

func (o TopRiskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopRiskInfo struct{}"
	}

	return strings.Join([]string{"TopRiskInfo", string(data)}, " ")
}
