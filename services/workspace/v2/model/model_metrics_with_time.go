package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricsWithTime 带时间的指标
type MetricsWithTime struct {

	// 时间
	Time *string `json:"time,omitempty"`

	// 指标值
	Metrics *[]Metric `json:"metrics,omitempty"`
}

func (o MetricsWithTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricsWithTime struct{}"
	}

	return strings.Join([]string{"MetricsWithTime", string(data)}, " ")
}
