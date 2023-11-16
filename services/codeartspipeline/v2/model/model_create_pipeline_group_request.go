package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePipelineGroupRequest Request Object
type CreatePipelineGroupRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	Body *PipelineGroupCreateDto `json:"body,omitempty"`
}

func (o CreatePipelineGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePipelineGroupRequest struct{}"
	}

	return strings.Join([]string{"CreatePipelineGroupRequest", string(data)}, " ")
}
