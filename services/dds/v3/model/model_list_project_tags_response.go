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
type ListProjectTagsResponse struct {
	// 标签列表。
	Tags           *[]QueryProjectTagItem `json:"tags,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListProjectTagsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListProjectTagsResponse", string(data)}, " ")
}
