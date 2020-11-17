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
type UpdateTopicReplicaResponse struct {
}

func (o UpdateTopicReplicaResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateTopicReplicaResponse", string(data)}, " ")
}
