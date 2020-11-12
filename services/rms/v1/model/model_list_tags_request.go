/*
 * RMS
 *
 * Resource Manager Api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListTagsRequest struct {
	TagKey *string `json:"tag_key,omitempty"`
	Marker *string `json:"marker,omitempty"`
	Limit  *int32  `json:"limit,omitempty"`
}

func (o ListTagsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListTagsRequest", string(data)}, " ")
}
