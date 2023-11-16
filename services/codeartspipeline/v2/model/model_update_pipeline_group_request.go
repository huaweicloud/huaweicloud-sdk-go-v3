package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePipelineGroupRequest Request Object
type UpdatePipelineGroupRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	Body *PipelineGroupUpdateDto `json:"body,omitempty"`
}

func (o UpdatePipelineGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePipelineGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdatePipelineGroupRequest", string(data)}, " ")
}
