package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModuleTemplateFieldDto struct {

	// 相关描述信息
	TemplateIds []string `json:"template_ids"`
}

func (o ModuleTemplateFieldDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModuleTemplateFieldDto struct{}"
	}

	return strings.Join([]string{"ModuleTemplateFieldDto", string(data)}, " ")
}
