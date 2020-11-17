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
type CreatePartitionRequest struct {
	ProjectId  string              `json:"project_id"`
	InstanceId string              `json:"instance_id"`
	Topic      string              `json:"topic"`
	Body       *CreatePartitionReq `json:"body,omitempty"`
}

func (o CreatePartitionRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreatePartitionRequest", string(data)}, " ")
}
