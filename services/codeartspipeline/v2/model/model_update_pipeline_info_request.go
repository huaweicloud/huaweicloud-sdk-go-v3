package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePipelineInfoRequest Request Object
type UpdatePipelineInfoRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 流水线ID
	PipelineId string `json:"pipeline_id"`

	// 微服务ID
	ComponentId *string `json:"componentId,omitempty"`

	Body *PipelineDto `json:"body,omitempty"`
}

func (o UpdatePipelineInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePipelineInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdatePipelineInfoRequest", string(data)}, " ")
}
