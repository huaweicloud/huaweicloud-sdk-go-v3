/*
 * Kafka
 *
 * Kafka Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowMessagesResponse struct {
	// 消息列表。
	Messages *[]ShowMessagesRespMessages `json:"messages,omitempty"`
	// 消息总数。
	MessagesCount *int32 `json:"messages_count,omitempty"`
	// 总页数。
	OffsetsCount *int32 `json:"offsets_count,omitempty"`
	// 当前页数。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowMessagesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowMessagesResponse", string(data)}, " ")
}
