package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 插件模板详细信息
type Templatespec struct {

	// 模板类型（helm，static）
	Type string `json:"type" xml:"type"`

	// 是否为必安装插件
	Require *bool `json:"require,omitempty" xml:"require"`

	// 模板所属分组
	Labels []string `json:"labels" xml:"labels"`

	// Logo图片地址
	LogoURL string `json:"logoURL" xml:"logoURL"`

	// 插件详情描述及使用说明
	ReadmeURL string `json:"readmeURL" xml:"readmeURL"`

	// 模板描述
	Description string `json:"description" xml:"description"`

	// 模板具体版本详情
	Versions []Versions `json:"versions" xml:"versions"`
}

func (o Templatespec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Templatespec struct{}"
	}

	return strings.Join([]string{"Templatespec", string(data)}, " ")
}
