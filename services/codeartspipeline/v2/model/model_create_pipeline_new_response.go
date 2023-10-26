package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePipelineNewResponse Response Object
type CreatePipelineNewResponse struct {

	// 流水线ID
	PipelineId     *string `json:"pipeline_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePipelineNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePipelineNewResponse struct{}"
	}

	return strings.Join([]string{"CreatePipelineNewResponse", string(data)}, " ")
}
