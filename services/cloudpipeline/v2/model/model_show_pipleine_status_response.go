package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowPipleineStatusResponse struct {

	// 流水线ID
	PipelineId *string `json:"pipeline_id,omitempty" xml:"pipeline_id"`

	// 流水线名称
	PipelineName *string `json:"pipeline_name,omitempty" xml:"pipeline_name"`

	// 执行人
	Executor *string `json:"executor,omitempty" xml:"executor"`

	// 流水线执行ID
	BuildId *string `json:"build_id,omitempty" xml:"build_id"`

	// 开始执行时间
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// 结束执行时间
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`

	// 流水线参数
	Parameters *[]PipelineParameter `json:"parameters,omitempty" xml:"parameters"`

	// 流水线执行情况
	States *[]PipelineStateStatus `json:"states,omitempty" xml:"states"`

	// 执行耗时
	ElapsedTime *string `json:"elapsed_time,omitempty" xml:"elapsed_time"`

	// 流水线运行状态
	Status *string `json:"status,omitempty" xml:"status"`

	// 流水线执行结果
	Outcome *string `json:"outcome,omitempty" xml:"outcome"`

	// 流水线详情页地址
	DetailUrl      *string `json:"detail_url,omitempty" xml:"detail_url"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPipleineStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipleineStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowPipleineStatusResponse", string(data)}, " ")
}
