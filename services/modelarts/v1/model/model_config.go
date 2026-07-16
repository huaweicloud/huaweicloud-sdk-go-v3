package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Config struct {

	// **参数解释**：自定义脚本内容（base64编码）或脚本绝对路径。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Script *string `json:"script,omitempty"`

	// **参数解释**：脚本类型。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - COMMAND：script中需要指定脚本内容（base64编码）。 - SCRIPT：script中需要指定脚本路径。  **默认取值**：SCRIPT。
	Type *string `json:"type,omitempty"`

	// **参数解释**：自定义脚本执行方式，同步或异步执行。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - BLOCK：同步 - ASYNC：异步  **默认取值**：ASYNC
	Mode *string `json:"mode,omitempty"`
}

func (o Config) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Config struct{}"
	}

	return strings.Join([]string{"Config", string(data)}, " ")
}
