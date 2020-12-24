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
type ListRestoreTimesResponse struct {
	// 可恢复时间段列表。
	RestoreTime    *[]GetRestoreTimeResponseRestoreTime `json:"restore_time,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ListRestoreTimesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListRestoreTimesResponse", string(data)}, " ")
}
