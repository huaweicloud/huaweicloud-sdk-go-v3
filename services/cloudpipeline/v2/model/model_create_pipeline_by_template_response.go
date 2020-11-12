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
type CreatePipelineByTemplateResponse struct {
	// 实例ID
	TaskId *string `json:"task_id,omitempty"`
}

func (o CreatePipelineByTemplateResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreatePipelineByTemplateResponse", string(data)}, " ")
}
