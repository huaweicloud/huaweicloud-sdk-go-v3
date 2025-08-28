package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LaunchTemplateOption struct {
	TemplateData *TemplateData `json:"template_data,omitempty"`

	// 模板名称
	Name string `json:"name"`

	// 模板描述
	Description *string `json:"description,omitempty"`

	// 初始第一个版本的版本描述
	VersionDescription *string `json:"version_description,omitempty"`
}

func (o LaunchTemplateOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LaunchTemplateOption struct{}"
	}

	return strings.Join([]string{"LaunchTemplateOption", string(data)}, " ")
}
