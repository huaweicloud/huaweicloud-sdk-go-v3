package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProjectTemplates struct {

	// 描述
	Description *string `json:"description,omitempty"`

	// 显示名
	DisplayName *string `json:"display_name,omitempty"`

	// 图标
	Logo *string `json:"logo,omitempty"`

	// 模板名
	Name *string `json:"name,omitempty"`

	// 路径
	Path *string `json:"path,omitempty"`

	// 项目类型
	ProjectType *string `json:"project_type,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`

	Source *SourceStorage `json:"source,omitempty"`

	// tags
	Tags *[]string `json:"tags,omitempty"`

	// 模板id
	TemplateId *int64 `json:"template_id,omitempty"`

	// cpu架构
	Arch *string `json:"arch,omitempty"`
}

func (o ProjectTemplates) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectTemplates struct{}"
	}

	return strings.Join([]string{"ProjectTemplates", string(data)}, " ")
}
