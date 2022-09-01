package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdatePipelineJobResponse struct {

	// 管道ID
	PipelineId *string `json:"pipeline_id,omitempty" xml:"pipeline_id"`

	// 管道状态
	PipelineState *string `json:"pipeline_state,omitempty" xml:"pipeline_state"`

	// 操作结果
	Status *string `json:"status,omitempty" xml:"status"`

	// 管道错误详情
	CheckInfo      map[string]interface{} `json:"check_info,omitempty" xml:"check_info"`
	HttpStatusCode int                    `json:"-"`
}

func (o UpdatePipelineJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePipelineJobResponse struct{}"
	}

	return strings.Join([]string{"UpdatePipelineJobResponse", string(data)}, " ")
}
