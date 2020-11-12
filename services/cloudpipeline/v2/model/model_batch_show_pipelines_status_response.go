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
type BatchShowPipelinesStatusResponse struct {
	// 流水线执行结果
	Result *string `json:"result,omitempty"`
	// 流水线执行状态
	Status *string `json:"status,omitempty"`
	Stages *Stages `json:"stages,omitempty"`
	// 执行人
	Executor *string `json:"executor,omitempty"`
	// 流水线名字
	PipelineName *string `json:"pipeline_name,omitempty"`
	// 流水线ID
	PipelineId *string `json:"pipeline_id,omitempty"`
	// 流水线详情页URL
	DetailUrl *string `json:"detail_url,omitempty"`
	// 流水线编辑页URL
	ModifyUrl *string `json:"modify_url,omitempty"`
	// 开始执行时间
	StartTime *string `json:"start_time,omitempty"`
	// 结束执行时间
	EndTime *string `json:"end_time,omitempty"`
}

func (o BatchShowPipelinesStatusResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchShowPipelinesStatusResponse", string(data)}, " ")
}
