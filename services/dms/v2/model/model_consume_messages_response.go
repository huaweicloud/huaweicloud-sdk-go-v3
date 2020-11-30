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
type ConsumeMessagesResponse struct {
	// 消息数组。
	Body           *[]ConsumeMessage `json:"body,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ConsumeMessagesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ConsumeMessagesResponse", string(data)}, " ")
}
