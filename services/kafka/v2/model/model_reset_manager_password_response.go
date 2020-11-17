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
type ResetManagerPasswordResponse struct {
}

func (o ResetManagerPasswordResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ResetManagerPasswordResponse", string(data)}, " ")
}
