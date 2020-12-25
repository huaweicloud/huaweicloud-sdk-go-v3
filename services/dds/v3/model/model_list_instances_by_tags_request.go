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
type ListInstancesByTagsRequest struct {
	Body *ListInstancesByTagsRequestBody `json:"body,omitempty"`
}

func (o ListInstancesByTagsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListInstancesByTagsRequest", string(data)}, " ")
}
