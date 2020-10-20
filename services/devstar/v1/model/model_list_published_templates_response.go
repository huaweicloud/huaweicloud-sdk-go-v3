/*
 * DevStar
 *
 * DevStar API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListPublishedTemplatesResponse struct {
	// 返回模板的数量
	Count *int32 `json:"count,omitempty"`
	// 返回模板的列表
	Templates *[]TemplateSimpleInfo `json:"templates,omitempty"`
}

func (o ListPublishedTemplatesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListPublishedTemplatesResponse", string(data)}, " ")
}
