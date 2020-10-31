/*
 * ServiceStage
 *
 * ServiceStage的API,包括应用管理和仓库授权管理
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListTemplatesResponse struct {
	// 模板列表。
	Templates *[]TemplateView `json:"templates,omitempty"`
}

func (o ListTemplatesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListTemplatesResponse", string(data)}, " ")
}
