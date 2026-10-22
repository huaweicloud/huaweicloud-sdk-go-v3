package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkerGroupConfig struct {

	// **参数解释**：工作组名称。 **约束限制**：已有工作组名称不支持修改，需和当前工作组名称保持一致；新增工作组该字段必填。 **取值范围**：长度为1~63个字符，支持字母、数字和中划线的组合。 **默认取值**：不涉及。
	Name string `json:"name"`

	// **参数解释**：节点规格。 **约束限制**：不涉及。 **取值范围**：支持英文字母、数字、小数点及连字符，长度为1~64个字符。 **默认取值**：不涉及。
	SpecCode string `json:"spec_code"`

	// **参数解释**：通算CPU大小，单位毫核。 **约束限制**：不涉及。 **取值范围**：0~2147483647。 **默认取值**：不涉及。
	Cpu *int32 `json:"cpu,omitempty"`

	// **参数解释**：通算内存大小，单位MB。 **约束限制**：不涉及。 **取值范围**：0~2147483647。 **默认取值**：不涉及。
	Memory *int32 `json:"memory,omitempty"`

	// **参数解释**：GPU卡数量。 **约束限制**：暂时只支持整数。 **取值范围**：0~2147483647。 **默认取值**：不涉及。
	Gpu *float64 `json:"gpu,omitempty"`

	// **参数解释**：NPU卡数量。 **约束限制**：暂时只支持整数。 **取值范围**：0~2147483647。 **默认取值**：不涉及。
	Npu *float64 `json:"npu,omitempty"`

	Image *ImageVersionInfo `json:"image,omitempty"`
}

func (o WorkerGroupConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkerGroupConfig struct{}"
	}

	return strings.Join([]string{"WorkerGroupConfig", string(data)}, " ")
}
