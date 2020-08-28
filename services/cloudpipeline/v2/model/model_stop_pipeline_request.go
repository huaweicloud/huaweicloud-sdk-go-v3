/*
 * devcloudpipeline
 *
 * devcloud pipeline api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type StopPipelineRequest struct {
	XLanguage  *string `json:"X-Language,omitempty"`
	PipelineId string  `json:"pipeline_id"`
	BuildId    *string `json:"build_id,omitempty"`
}

func (o StopPipelineRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"StopPipelineRequest", string(data)}, " ")
}
