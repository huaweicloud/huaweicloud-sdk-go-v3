/*
 * RDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RestoreToExistingInstanceResponse struct {
	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestoreToExistingInstanceResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RestoreToExistingInstanceResponse", string(data)}, " ")
}
