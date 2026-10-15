package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobRuntimeEnv **参数解释**：作业运行环境配置。  **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type JobRuntimeEnv struct {

	// **参数解释**：作业运行工作目录。  **约束限制**：不涉及。 **取值范围**：长度为1~1024的任意字符组合。 **默认取值**：不涉及。
	WorkingDir string `json:"working_dir"`

	// **参数解释**：作业依赖的zip包路径。  **约束限制**：不涉及。 **取值范围**：长度为1~1024的任意字符组合。 **默认取值**：不涉及。
	DependencyDir *[]string `json:"dependency_dir,omitempty"`

	// **参数解释**：作业运行Dataframe版本号。  **约束限制**：不涉及。 **取值范围**：长度为1~1024的任意字符组合。 **默认取值**：不涉及。
	PythonVersion *string `json:"python_version,omitempty"`

	// **参数解释**：作业运行环境变量。  **约束限制**：不涉及。 **取值范围**：长度为0~1024。 **默认取值**：不涉及。
	EnvVars map[string]string `json:"env_vars,omitempty"`

	// **参数解释**：作业运行依赖的软件包列表。  **约束限制**：不涉及。 **取值范围**：长度为0~1024。 **默认取值**：不涉及。
	Pip *[]string `json:"pip,omitempty"`
}

func (o JobRuntimeEnv) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobRuntimeEnv struct{}"
	}

	return strings.Join([]string{"JobRuntimeEnv", string(data)}, " ")
}
