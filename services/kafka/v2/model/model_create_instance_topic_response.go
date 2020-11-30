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
type CreateInstanceTopicResponse struct {
	// topic名称。
	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInstanceTopicResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateInstanceTopicResponse", string(data)}, " ")
}
