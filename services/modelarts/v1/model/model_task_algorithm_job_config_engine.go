package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskAlgorithmJobConfigEngine **参数解释**：算法的引擎。 **约束限制**：不涉及。
type TaskAlgorithmJobConfigEngine struct {

	// **参数解释**：算法选择的引擎规格ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	EngineId *string `json:"engine_id,omitempty"`

	// **参数解释**：算法选择的引擎规格名称。 **约束限制**：若填入engine_id则无需填写。 **取值范围**：不涉及。 **默认取值**：不涉及。
	EngineName *string `json:"engine_name,omitempty"`

	// **参数解释**：算法选择的引擎规格版本。 **约束限制**：若填入engine_id则无需填写。 **取值范围**：不涉及。 **默认取值**：不涉及。
	EngineVersion *string `json:"engine_version,omitempty"`

	// **参数解释**：算法选择的自定义镜像地址。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ImageUrl *string `json:"image_url,omitempty"`

	// **参数解释**：容器镜像启动用户，默认为1000，仅自定义镜像场景下支持配置。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	RunUser *string `json:"run_user,omitempty"`
}

func (o TaskAlgorithmJobConfigEngine) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskAlgorithmJobConfigEngine struct{}"
	}

	return strings.Join([]string{"TaskAlgorithmJobConfigEngine", string(data)}, " ")
}
