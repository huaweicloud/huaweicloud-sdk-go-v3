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
type ListResourceInstancesRequest struct {
	ResourceType string                   `json:"resource_type"`
	Body         *ListInstanceRequestBody `json:"body,omitempty"`
}

func (o ListResourceInstancesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListResourceInstancesRequest", string(data)}, " ")
}
