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
type DeleteSinkTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSinkTaskResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteSinkTaskResponse", string(data)}, " ")
}
