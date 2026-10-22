package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServeRuntimeEnv **参数解释**：Serve 应用运行时环境配置。
type ServeRuntimeEnv struct {

	// **参数解释**：工作目录，如 file:///home/service/works/code/.../xxx.zip。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	WorkingDir string `json:"working_dir"`

	// **参数解释**：环境变量映射。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	EnvVars map[string]string `json:"env_vars,omitempty"`
}

func (o ServeRuntimeEnv) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServeRuntimeEnv struct{}"
	}

	return strings.Join([]string{"ServeRuntimeEnv", string(data)}, " ")
}
