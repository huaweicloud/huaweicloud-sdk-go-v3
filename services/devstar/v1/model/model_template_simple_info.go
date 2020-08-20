/*
 * Devstar
 *
 * Devstar API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type TemplateSimpleInfo struct {
	// 模板id
	Id *string `json:"id,omitempty"`
	// 模板名
	Title *string `json:"title,omitempty"`
	// 模板描述
	Description *string `json:"description,omitempty"`
}

func (o TemplateSimpleInfo) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"TemplateSimpleInfo", string(data)}, " ")
}
