package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuntimeEnv **参数解释**：作业的运行时环境配置。 **约束限制**：不涉及。 **默认取值**：不涉及。
type RuntimeEnv struct {

	// **参数解释**：代码将在其中运行的工作目录。 **约束限制**：必须是远程URI，如OBS路径。 **取值范围**：最小长度1，最大长度2048。 **默认取值**：不涉及。
	WorkingDir *string `json:"working_dir,omitempty"`

	// **参数解释**：要安装的pip软件包列表。 **约束限制**：最小个数0，最大个数1023。
	Pip *[]string `json:"pip,omitempty"`

	// **参数解释**：要设置的环境变量。key为环境变量名称，value为环境变量值，均为String类型。 **约束限制**：不涉及。
	EnvVars map[string]string `json:"env_vars,omitempty"`

	// **参数解释**：运行环境的配置。key为配置项名称，value为配置项值。 **约束限制**：不涉及。
	Config map[string]interface{} `json:"config,omitempty"`
}

func (o RuntimeEnv) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuntimeEnv struct{}"
	}

	return strings.Join([]string{"RuntimeEnv", string(data)}, " ")
}
