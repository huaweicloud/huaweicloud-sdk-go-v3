package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkFlowMetricResponse Response Object
type ShowWorkFlowMetricResponse struct {

	// 执行次数
	Count *[]SlaReportsValue `json:"count,omitempty"`

	// 平均时延，单位毫秒
	Duration *[]SlaReportsValue `json:"duration,omitempty"`

	// 错误次数
	FailCount *[]SlaReportsValue `json:"fail_count,omitempty"`

	// 运行中数量
	RunningCount   *[]SlaReportsValue `json:"running_count,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowWorkFlowMetricResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkFlowMetricResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkFlowMetricResponse", string(data)}, " ")
}
