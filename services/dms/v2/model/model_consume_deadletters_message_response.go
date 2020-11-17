/*
 * DMS
 *
 * DMS Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ConsumeDeadlettersMessageResponse struct {
	// 消息数组。
	Body *[]ConsumeDeadlettersMessage `json:"body,omitempty"`
}

func (o ConsumeDeadlettersMessageResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ConsumeDeadlettersMessageResponse", string(data)}, " ")
}
