package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemovePipelineRequest Request Object
type RemovePipelineRequest struct {

	// 要删除的流水线ID
	PipelineId string `json:"pipeline_id"`
}

func (o RemovePipelineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemovePipelineRequest struct{}"
	}

	return strings.Join([]string{"RemovePipelineRequest", string(data)}, " ")
}
