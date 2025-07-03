package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateVersion struct {
	TemplateData *TemplateData `json:"template_data,omitempty"`

	// 版本号
	VersionNumber *int32 `json:"version_number,omitempty"`

	// 版本id
	VersionId *string `json:"version_id,omitempty"`

	// 版本描述
	VersionDescription *string `json:"version_description,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 模板id
	LaunchTemplateId *string `json:"launch_template_id,omitempty"`
}

func (o TemplateVersion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateVersion struct{}"
	}

	return strings.Join([]string{"TemplateVersion", string(data)}, " ")
}
