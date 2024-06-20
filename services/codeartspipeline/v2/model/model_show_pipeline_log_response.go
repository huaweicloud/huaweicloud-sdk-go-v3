package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPipelineLogResponse Response Object
type ShowPipelineLogResponse struct {

	// 是否有更多日志
	HasMore *bool `json:"has_more,omitempty"`

	// 查询日志结束偏移。填入请求体end_offset字段，用于查询下一页日志。
	EndOffset *string `json:"end_offset,omitempty"`

	// 查询日志起始偏移。填入请求体start_offset字段，用于查询下一页日志。
	StartOffset *string `json:"start_offset,omitempty"`

	// 日志内容
	Log *string `json:"log,omitempty"`

	// 日志存储位置
	Location *string `json:"location,omitempty"`

	// 所属步骤ID
	StepRunId      *string `json:"step_run_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPipelineLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipelineLogResponse struct{}"
	}

	return strings.Join([]string{"ShowPipelineLogResponse", string(data)}, " ")
}
