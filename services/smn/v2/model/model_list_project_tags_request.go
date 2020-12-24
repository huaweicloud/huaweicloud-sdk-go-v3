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
type ListProjectTagsRequest struct {
	ResourceType string `json:"resource_type"`
}

func (o ListProjectTagsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListProjectTagsRequest", string(data)}, " ")
}
