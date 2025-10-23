package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateEntity struct {

	// 模板ID
	TemplateId string `json:"template_id"`

	// 模板名称
	TemplateName string `json:"template_name"`

	TemplateType *TemplateTypeEnum `json:"template_type"`

	// 模板创建时间
	CreateTime string `json:"create_time"`

	// 最后更新时间
	LastUpdateTime string `json:"last_update_time"`

	// 模板具体的内容项，例如：存储库-使用率，存储库-摘要，任务-趋势，任务-摘要，任务-详情
	TemplateContents []TemplateItemEnum `json:"template_contents"`
}

func (o TemplateEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateEntity struct{}"
	}

	return strings.Join([]string{"TemplateEntity", string(data)}, " ")
}
