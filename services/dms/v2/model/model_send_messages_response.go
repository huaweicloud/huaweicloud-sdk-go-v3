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
type SendMessagesResponse struct {
	// 消息列表。
	Messages *[]SendMessagesRespMessages `json:"messages,omitempty"`
}

func (o SendMessagesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"SendMessagesResponse", string(data)}, " ")
}
