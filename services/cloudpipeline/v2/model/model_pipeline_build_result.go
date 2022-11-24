package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineBuildResult struct {

	// 执行ID
	BuildId string `json:"build_id"`

	// 运行耗时
	ElapseTime *string `json:"elapse_time,omitempty"`

	// 执行结束时间
	EndTime string `json:"end_time"`

	// 运行结果
	Outcome string `json:"outcome"`

	// 流水线id
	PipelineId string `json:"pipeline_id"`

	// 流水线名称
	PipelineName string `json:"pipeline_name"`

	// 执行开始时间
	StartTime string `json:"start_time"`

	// 运行状态
	Status string `json:"status"`
}

func (o PipelineBuildResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineBuildResult struct{}"
	}

	return strings.Join([]string{"PipelineBuildResult", string(data)}, " ")
}
