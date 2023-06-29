package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunPipelineResponse Response Object
type RunPipelineResponse struct {
	PipelineRunId  *string `json:"pipeline_run_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunPipelineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunPipelineResponse struct{}"
	}

	return strings.Join([]string{"RunPipelineResponse", string(data)}, " ")
}
