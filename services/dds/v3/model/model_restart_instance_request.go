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
type RestartInstanceRequest struct {
	InstanceId string                      `json:"instance_id"`
	Body       *RestartInstanceRequestBody `json:"body,omitempty"`
}

func (o RestartInstanceRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RestartInstanceRequest", string(data)}, " ")
}
