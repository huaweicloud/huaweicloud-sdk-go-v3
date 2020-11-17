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
type DeleteBackgroundTaskResponse struct {
}

func (o DeleteBackgroundTaskResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteBackgroundTaskResponse", string(data)}, " ")
}
