package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DasMetricInfo DAS指标信息
type DasMetricInfo struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 阈值索引
	ThresholdIndex *int32 `json:"threshold_index,omitempty"`

	// 实例状态
	InstanceStatus *string `json:"instance_status,omitempty"`

	// 指标采集时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 指标数据
	Metrics *interface{} `json:"metrics,omitempty"`

	// 阈值指标数据
	ThresholdMetrics *interface{} `json:"threshold_metrics,omitempty"`
}

func (o DasMetricInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DasMetricInfo struct{}"
	}

	return strings.Join([]string{"DasMetricInfo", string(data)}, " ")
}
