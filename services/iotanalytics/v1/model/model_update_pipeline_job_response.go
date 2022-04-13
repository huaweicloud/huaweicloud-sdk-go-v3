package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdatePipelineJobResponse struct {
	// 管道ID

	PipelineId *string `json:"pipeline_id,omitempty"`
	// 管道状态

	PipelineState *string `json:"pipeline_state,omitempty"`
	// 操作结果

	Status *string `json:"status,omitempty"`
	// 管道错误详情

	CheckInfo      map[string]interface{} `json:"check_info,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o UpdatePipelineJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePipelineJobResponse struct{}"
	}

	return strings.Join([]string{"UpdatePipelineJobResponse", string(data)}, " ")
}
