package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProjectTemplates struct {

	// 描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 显示名
	DisplayName *string `json:"display_name,omitempty" xml:"display_name"`

	// 图标
	Logo *string `json:"logo,omitempty" xml:"logo"`

	// 模板名
	Name *string `json:"name,omitempty" xml:"name"`

	// 路径
	Path *string `json:"path,omitempty" xml:"path"`

	// 项目类型
	ProjectType *string `json:"project_type,omitempty" xml:"project_type"`

	// 区域
	Region *string `json:"region,omitempty" xml:"region"`

	Source *SourceStorage `json:"source,omitempty" xml:"source"`

	// tags
	Tags *[]string `json:"tags,omitempty" xml:"tags"`

	// 模板id
	TemplateId *int64 `json:"template_id,omitempty" xml:"template_id"`

	// cpu架构
	Arch *string `json:"arch,omitempty" xml:"arch"`
}

func (o ProjectTemplates) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectTemplates struct{}"
	}

	return strings.Join([]string{"ProjectTemplates", string(data)}, " ")
}
