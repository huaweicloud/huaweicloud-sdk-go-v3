/*
 * smn
 *
 * SMN Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateApplicationEndpointRequest struct {
	EndpointUrn string                                `json:"endpoint_urn"`
	Body        *UpdateApplicationEndpointRequestBody `json:"body,omitempty"`
}

func (o UpdateApplicationEndpointRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateApplicationEndpointRequest", string(data)}, " ")
}
