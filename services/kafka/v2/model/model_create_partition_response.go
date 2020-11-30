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
type CreatePartitionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreatePartitionResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreatePartitionResponse", string(data)}, " ")
}
