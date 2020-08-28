/*
 * devcloudpipeline
 *
 * devcloud pipeline api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RegisterAgentRequest struct {
	Body *SlaveRegister `json:"body,omitempty"`
}

func (o RegisterAgentRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RegisterAgentRequest", string(data)}, " ")
}
