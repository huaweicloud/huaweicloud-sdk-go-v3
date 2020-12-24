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
type ListSlowLogsResponse struct {
	SlowLogList *[]SlowLog `json:"slow_log_list,omitempty"`
	// 总记录数。
	TotalRecord    *int32 `json:"total_record,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSlowLogsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListSlowLogsResponse", string(data)}, " ")
}
