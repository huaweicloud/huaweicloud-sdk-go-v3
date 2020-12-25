/*
 * DDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateInstanceNameRequest struct {
	InstanceId string                 `json:"instance_id"`
	Body       *UpdateNameRequestBody `json:"body,omitempty"`
}

func (o UpdateInstanceNameRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateInstanceNameRequest", string(data)}, " ")
}
