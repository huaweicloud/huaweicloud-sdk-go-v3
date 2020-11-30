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
type RunDevstarTemplateJobResponse struct {
	// 任务id
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunDevstarTemplateJobResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RunDevstarTemplateJobResponse", string(data)}, " ")
}
