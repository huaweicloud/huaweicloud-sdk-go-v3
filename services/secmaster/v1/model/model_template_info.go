package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TemplateInfo 插件配置模板信息
type TemplateInfo struct {

	// 模板id
	Id *string `json:"id,omitempty"`

	// 插件id
	ComponentId *string `json:"component_id,omitempty"`

	// 插件配置模板名称
	TemplateName *string `json:"template_name,omitempty"`

	// 插件action的配置内容
	TaskConfig *string `json:"task_config,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`
}

func (o TemplateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateInfo struct{}"
	}

	return strings.Join([]string{"TemplateInfo", string(data)}, " ")
}
