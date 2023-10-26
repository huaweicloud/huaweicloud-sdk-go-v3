package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePipelineNewRequest Request Object
type CreatePipelineNewRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 组件ID
	ComponentId *string `json:"component_id,omitempty"`

	Body *PipelineDto `json:"body,omitempty"`
}

func (o CreatePipelineNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePipelineNewRequest struct{}"
	}

	return strings.Join([]string{"CreatePipelineNewRequest", string(data)}, " ")
}
