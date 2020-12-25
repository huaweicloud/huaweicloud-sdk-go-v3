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

// Response Object
type ListInstanceTagsResponse struct {
	// 标签列表。
	Tags           *[]QueryResourceTagItem `json:"tags,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListInstanceTagsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListInstanceTagsResponse", string(data)}, " ")
}
