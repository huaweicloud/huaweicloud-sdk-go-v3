package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchMovePipelineToGroupRequest Request Object
type BatchMovePipelineToGroupRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	Body *PipelineGroupBindDto `json:"body,omitempty"`
}

func (o BatchMovePipelineToGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchMovePipelineToGroupRequest struct{}"
	}

	return strings.Join([]string{"BatchMovePipelineToGroupRequest", string(data)}, " ")
}
