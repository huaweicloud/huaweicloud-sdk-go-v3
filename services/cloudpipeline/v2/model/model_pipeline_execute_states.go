package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 流水线执行结果
type PipelineExecuteStates struct {

	// 流水线执行结果
	Result string `json:"result" xml:"result"`

	// 流水线执行状态
	Status string `json:"status" xml:"status"`

	// 阶段执行情况
	Stages []Stages `json:"stages" xml:"stages"`

	// 执行人
	Executor string `json:"executor" xml:"executor"`

	// 流水线名字
	PipelineName string `json:"pipeline_name" xml:"pipeline_name"`

	// 流水线ID
	PipelineId string `json:"pipeline_id" xml:"pipeline_id"`

	// 流水线详情页URL
	DetailUrl string `json:"detail_url" xml:"detail_url"`

	// 流水线编辑页URL
	ModifyUrl string `json:"modify_url" xml:"modify_url"`

	// 开始执行时间
	StartTime string `json:"start_time" xml:"start_time"`

	// 结束执行时间
	EndTime string `json:"end_time" xml:"end_time"`
}

func (o PipelineExecuteStates) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineExecuteStates struct{}"
	}

	return strings.Join([]string{"PipelineExecuteStates", string(data)}, " ")
}
