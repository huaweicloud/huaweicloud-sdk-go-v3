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

// Request Object
type ShowPartitionEndMessageRequest struct {
	ProjectId  string `json:"project_id"`
	InstanceId string `json:"instance_id"`
	Topic      string `json:"topic"`
	Partition  int32  `json:"partition"`
}

func (o ShowPartitionEndMessageRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowPartitionEndMessageRequest", string(data)}, " ")
}
