/*
 * DDS
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
type RestartInstanceResponse struct {
	// 工作流ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestartInstanceResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RestartInstanceResponse", string(data)}, " ")
}
