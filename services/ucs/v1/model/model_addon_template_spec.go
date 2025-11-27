package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddonTemplateSpec struct {

	// 插件的安装类型，支持helm安装或static安装
	Type *string `json:"type,omitempty"`

	// 该插件是否为必装
	Require *bool `json:"require,omitempty"`

	// 插件的标签列表
	Labels *[]string `json:"labels,omitempty"`

	// Logo链接
	LogoURL *string `json:"logoURL,omitempty"`

	// README文档链接
	ReadmeURL *string `json:"readmeURL,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 插件的版本列表
	Versions *[]AddonVersion `json:"versions,omitempty"`
}

func (o AddonTemplateSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddonTemplateSpec struct{}"
	}

	return strings.Join([]string{"AddonTemplateSpec", string(data)}, " ")
}
