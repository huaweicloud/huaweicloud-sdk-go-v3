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
type UpdateInstanceRequest struct {
	ProjectId  string             `json:"project_id"`
	InstanceId string             `json:"instance_id"`
	Body       *UpdateInstanceReq `json:"body,omitempty"`
}

func (o UpdateInstanceRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateInstanceRequest", string(data)}, " ")
}
