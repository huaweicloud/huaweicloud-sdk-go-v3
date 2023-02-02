package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunPipelineRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 流水线ID
	PipelineId string `json:"pipeline_id"`

	// 语言类型 中文:zh-cn 英文:en-us，默认en-us
	XLanguage *string `json:"X-Language,omitempty"`

	Body *RunPipelineDto `json:"body,omitempty"`
}

func (o RunPipelineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunPipelineRequest struct{}"
	}

	return strings.Join([]string{"RunPipelineRequest", string(data)}, " ")
}
