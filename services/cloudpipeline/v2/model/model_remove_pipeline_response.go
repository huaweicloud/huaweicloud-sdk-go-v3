/*
 * CloudPipeline
 *
 * devcloud pipeline api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RemovePipelineResponse struct {
	// 流水线ID
	PipelineId *string `json:"pipeline_id,omitempty"`
	// 流水线名字
	PipelineName *string `json:"pipeline_name,omitempty"`
}

func (o RemovePipelineResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RemovePipelineResponse", string(data)}, " ")
}
