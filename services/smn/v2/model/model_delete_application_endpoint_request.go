/*
 * SMN
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
type DeleteApplicationEndpointRequest struct {
	EndpointUrn string `json:"endpoint_urn"`
}

func (o DeleteApplicationEndpointRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteApplicationEndpointRequest", string(data)}, " ")
}
