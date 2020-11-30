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
type UpdateInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateInstanceResponse", string(data)}, " ")
}
