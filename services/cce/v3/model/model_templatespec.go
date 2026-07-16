package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Templatespec 插件模板详细信息
type Templatespec struct {

	// **参数解释：** 插件模板类型 **约束限制：** 不涉及 **取值范围：** - helm: 表示使用Helm包进行部署的模板类型 - static: 表示静态模板类型  **默认取值：** 不涉及
	Type string `json:"type"`

	// 是否为必安装插件
	Require *bool `json:"require,omitempty"`

	// 模板所属分组
	Labels []string `json:"labels"`

	// Logo图片地址
	LogoURL string `json:"logoURL"`

	// 插件详情描述及使用说明
	ReadmeURL string `json:"readmeURL"`

	// 模板描述
	Description string `json:"description"`

	// 模板具体版本详情
	Versions []Versions `json:"versions"`
}

func (o Templatespec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Templatespec struct{}"
	}

	return strings.Join([]string{"Templatespec", string(data)}, " ")
}
