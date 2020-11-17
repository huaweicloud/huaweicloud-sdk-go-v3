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
type ResetPasswordRequest struct {
	ProjectId  string            `json:"project_id"`
	InstanceId string            `json:"instance_id"`
	Body       *ResetPasswordReq `json:"body,omitempty"`
}

func (o ResetPasswordRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ResetPasswordRequest", string(data)}, " ")
}
