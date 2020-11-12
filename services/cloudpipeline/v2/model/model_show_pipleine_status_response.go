/*
 * CloudPipeline
 *
 * devcloud pipeline api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowPipleineStatusResponse struct {
	// 流水线ID
	PipelineId *string `json:"pipeline_id,omitempty"`
	// 流水线名称
	PipelineName *string `json:"pipeline_name,omitempty"`
	// 执行人
	Executor *string `json:"executor,omitempty"`
	// 流水线执行ID
	BuildId *string `json:"build_id,omitempty"`
	// 开始执行时间
	StartTime *string `json:"start_time,omitempty"`
	// 结束执行时间
	EndTime    *string              `json:"end_time,omitempty"`
	Parameters *PipelineParameter   `json:"parameters,omitempty"`
	States     *PipelineStateStatus `json:"states,omitempty"`
	// 执行耗时
	ElapsedTime *string `json:"elapsed_time,omitempty"`
	// 流水线运行状态
	Status *string `json:"status,omitempty"`
	// 流水线执行结果
	Outcome *string `json:"outcome,omitempty"`
	// 流水线详情页地址
	DetailUrl *string `json:"detail_url,omitempty"`
}

func (o ShowPipleineStatusResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowPipleineStatusResponse", string(data)}, " ")
}
