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

// Request Object
type StartPipelineRequest struct {
	XLanguage  *string `json:"X-Language,omitempty"`
	PipelineId string  `json:"pipeline_id"`
}

func (o StartPipelineRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"StartPipelineRequest", string(data)}, " ")
}
