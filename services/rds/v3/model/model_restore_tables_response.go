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
type RestoreTablesResponse struct {
	// 任务ID。
	JobId          *string `json:"jobId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestoreTablesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RestoreTablesResponse", string(data)}, " ")
}
