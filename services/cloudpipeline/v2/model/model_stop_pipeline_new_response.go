package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type StopPipelineNewResponse struct {
	// 流水线ID

	PipelineId *string `json:"pipeline_id,omitempty"`
	// 流水线名字

	PipelineName   *string `json:"pipeline_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopPipelineNewResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StopPipelineNewResponse struct{}"
	}

	return strings.Join([]string{"StopPipelineNewResponse", string(data)}, " ")
}
