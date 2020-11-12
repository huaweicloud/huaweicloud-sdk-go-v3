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

// Response Object
type ListTagsResponse struct {
	// 标签列表
	Value    *[]TagEntity `json:"value,omitempty"`
	PageInfo *PageInfo    `json:"page_info,omitempty"`
}

func (o ListTagsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListTagsResponse", string(data)}, " ")
}
