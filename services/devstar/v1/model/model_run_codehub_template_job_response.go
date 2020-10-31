/*
 * DevStar
 *
 * DevStar API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RunCodehubTemplateJobResponse struct {
	// 任务id
	JobId *string `json:"job_id,omitempty"`
}

func (o RunCodehubTemplateJobResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RunCodehubTemplateJobResponse", string(data)}, " ")
}
