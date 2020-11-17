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
type ResetManagerPasswordRequest struct {
	ProjectId  string                   `json:"project_id"`
	InstanceId string                   `json:"instance_id"`
	Body       *ResetManagerPasswordReq `json:"body,omitempty"`
}

func (o ResetManagerPasswordRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ResetManagerPasswordRequest", string(data)}, " ")
}
