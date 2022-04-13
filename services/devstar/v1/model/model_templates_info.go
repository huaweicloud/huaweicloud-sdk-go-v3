package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplatesInfo struct {
	// 模板ID列表。

	TemplateIds []string `json:"template_ids"`
	// 平台来源： - 0：codelabs - 1：devstar

	PlatformSource int32 `json:"platform_source"`
}

func (o TemplatesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplatesInfo struct{}"
	}

	return strings.Join([]string{"TemplatesInfo", string(data)}, " ")
}
