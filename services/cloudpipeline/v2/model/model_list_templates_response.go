/*
 * devcloudpipeline
 *
 * devcloud pipeline api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListTemplatesResponse struct {
	// 总数
	Total *int32 `json:"total,omitempty"`
	// 页码数
	PageNumber *int32 `json:"page_number,omitempty"`
	// 每页显示数
	PageSize *int32        `json:"page_size,omitempty"`
	Content  *TemplateView `json:"content,omitempty"`
}

func (o ListTemplatesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListTemplatesResponse", string(data)}, " ")
}
