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
type ResetMessageOffsetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetMessageOffsetResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ResetMessageOffsetResponse", string(data)}, " ")
}
