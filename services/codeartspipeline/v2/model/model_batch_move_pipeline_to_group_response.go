package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchMovePipelineToGroupResponse Response Object
type BatchMovePipelineToGroupResponse struct {
	Body           *[]PipelineMoveToGroupResponseVo `json:"body,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o BatchMovePipelineToGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchMovePipelineToGroupResponse struct{}"
	}

	return strings.Join([]string{"BatchMovePipelineToGroupResponse", string(data)}, " ")
}
