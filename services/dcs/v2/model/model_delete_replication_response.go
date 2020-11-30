/*
 * DCS
 *
 * DCS V2版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteReplicationResponse struct {
	// 删除副本的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteReplicationResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteReplicationResponse", string(data)}, " ")
}
