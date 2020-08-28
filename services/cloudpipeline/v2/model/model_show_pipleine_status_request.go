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
type ShowPipleineStatusRequest struct {
	XLanguage  *string `json:"X-Language,omitempty"`
	PipelineId string  `json:"pipeline_id"`
	BuildId    *string `json:"build_id,omitempty"`
}

func (o ShowPipleineStatusRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowPipleineStatusRequest", string(data)}, " ")
}
