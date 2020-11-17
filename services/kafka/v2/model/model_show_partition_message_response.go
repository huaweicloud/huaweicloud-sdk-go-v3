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
type ShowPartitionMessageResponse struct {
	// 消息列表。
	Message *[]ShowPartitionMessageRespMessage `json:"message,omitempty"`
}

func (o ShowPartitionMessageResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowPartitionMessageResponse", string(data)}, " ")
}
