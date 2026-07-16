package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Template 插件实例的模板信息。
type Template struct {

	// **参数解释**：待安装插件模板名称，如log-agent。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Name string `json:"name"`

	// **参数解释**：待安装、升级插件的版本号。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释**：插件模板安装参数（各插件不同），升级插件时需要填写全量安装参数，未填写参数将使用插件模板中的默认值，当前插件安装参数可通过查询插件实例接口获取。 **约束限制**：不涉及。
	Inputs map[string]interface{} `json:"inputs,omitempty"`
}

func (o Template) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Template struct{}"
	}

	return strings.Join([]string{"Template", string(data)}, " ")
}
