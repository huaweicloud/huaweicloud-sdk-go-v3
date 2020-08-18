/*
 * Devstar
 *
 * Devstar API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Response Object
type RunTemplateJobV2Response struct {
	// 任务id
	JobId *string `json:"job_id,omitempty"`
}

func (o RunTemplateJobV2Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RunTemplateJobV2Response", string(data)}, " ")
}
