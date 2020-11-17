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
type BatchDeleteInstanceTopicResponse struct {
	// Topic列表。
	Topics *[]BatchDeleteInstanceTopicRespTopics `json:"topics,omitempty"`
}

func (o BatchDeleteInstanceTopicResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchDeleteInstanceTopicResponse", string(data)}, " ")
}
