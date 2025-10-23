package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTemplateBody 更新模板配置
type UpdateTemplateBody struct {

	// 模板内容列表
	TemplateContents []TemplateItemEnum `json:"template_contents"`
}

func (o UpdateTemplateBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTemplateBody struct{}"
	}

	return strings.Join([]string{"UpdateTemplateBody", string(data)}, " ")
}
