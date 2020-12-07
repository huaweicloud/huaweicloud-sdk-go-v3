/*
 * smn
 *
 * SMN Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type UpdateMessageTemplateRequestBody struct {
	// 模板内容。
	Content string `json:"content"`
}

func (o UpdateMessageTemplateRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateMessageTemplateRequestBody", string(data)}, " ")
}
