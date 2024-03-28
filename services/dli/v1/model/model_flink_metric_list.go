package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkMetricList 所有作业监控信息。
type FlinkMetricList struct {

	// 作业ID。
	JobId *int64 `json:"job_id,omitempty"`

	Metrics *Metric `json:"metrics,omitempty"`
}

func (o FlinkMetricList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkMetricList struct{}"
	}

	return strings.Join([]string{"FlinkMetricList", string(data)}, " ")
}
