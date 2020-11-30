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
type UpdateInstanceAutoCreateTopicResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceAutoCreateTopicResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateInstanceAutoCreateTopicResponse", string(data)}, " ")
}
