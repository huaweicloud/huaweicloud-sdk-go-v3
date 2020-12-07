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
type PublishAppMessageRequest struct {
	EndpointUrn string                        `json:"endpoint_urn"`
	Body        *PublishAppMessageRequestBody `json:"body,omitempty"`
}

func (o PublishAppMessageRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"PublishAppMessageRequest", string(data)}, " ")
}
