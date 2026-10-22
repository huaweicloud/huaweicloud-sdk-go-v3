package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServeApplication **参数解释**：Serve 应用配置。
type ServeApplication struct {

	// **参数解释**：应用名（L7 路由/计费维度）。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Name string `json:"name"`

	// **参数解释**：Python 导入路径（模块:对象），如 voc_deepseek_server_v4:app。 **约束限制**：不能包含空格。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ImportPath string `json:"import_path"`

	// **参数解释**：路由前缀，端点内不可重复。 **约束限制**：必须以 / 开头，不能包含空格，且同一端点内多个应用的 route_prefix 不可重复。 **取值范围**：不涉及。 **默认取值**：不涉及。
	RoutePrefix string `json:"route_prefix"`

	RuntimeEnv *ServeRuntimeEnv `json:"runtime_env"`

	Deployments []ServeDeployment `json:"deployments"`
}

func (o ServeApplication) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServeApplication struct{}"
	}

	return strings.Join([]string{"ServeApplication", string(data)}, " ")
}
