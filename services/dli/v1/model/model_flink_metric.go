package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkMetric 作业审计日志信息。
type FlinkMetric struct {

	// 所有作业监控信息。
	Jobs *[]FlinkMetricList `json:"jobs,omitempty"`
}

func (o FlinkMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkMetric struct{}"
	}

	return strings.Join([]string{"FlinkMetric", string(data)}, " ")
}
