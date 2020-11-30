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
type DeleteInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteInstanceResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteInstanceResponse", string(data)}, " ")
}
