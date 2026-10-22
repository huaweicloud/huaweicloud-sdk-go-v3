package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContainerSoftware **参数解释**：镜像包包含的软件信息，帮助用户参考使用。
type ContainerSoftware struct {

	// **参数解释**：软件名称。 **约束限制**：不涉及。 **取值范围**：支持英文字母、数字、小数点、下划线和连字符，2~32个字符。 **默认取值**：不涉及。
	Name string `json:"name"`

	// **参数解释**：软件版本。 **约束限制**：不涉及。 **取值范围**：支持英文字母、数字、小数点、下划线和连字符，2~32个字符。 **默认取值**：不涉及。
	Version string `json:"version"`
}

func (o ContainerSoftware) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerSoftware struct{}"
	}

	return strings.Join([]string{"ContainerSoftware", string(data)}, " ")
}
