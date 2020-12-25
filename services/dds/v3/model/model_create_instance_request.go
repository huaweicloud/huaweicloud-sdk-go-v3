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
type CreateInstanceRequest struct {
	Body *CreateInstanceRequestBody `json:"body,omitempty"`
}

func (o CreateInstanceRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateInstanceRequest", string(data)}, " ")
}
