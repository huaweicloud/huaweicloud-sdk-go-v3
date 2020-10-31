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
type ListTemplatesV2Response struct {
	// 返回模板的数量
	Count *int32 `json:"count,omitempty"`
	// 返回模板的列表
	Templates *[]TemplateInfo `json:"templates,omitempty"`
}

func (o ListTemplatesV2Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListTemplatesV2Response", string(data)}, " ")
}
