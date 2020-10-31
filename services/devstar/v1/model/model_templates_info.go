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

type TemplatesInfo struct {
	// 模板ID列表
	TemplateIds []string `json:"template_ids"`
	// 平台来源（0:codelabs、1:devstar）
	PlatformSource int32 `json:"platform_source"`
}

func (o TemplatesInfo) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"TemplatesInfo", string(data)}, " ")
}
