package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoListLifecycleManagedModelUpdateLifecycleStateDto struct {

	// **参数解释：**  迭代版本。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Params *[]LifecycleManagedModelUpdateLifecycleStateDto `json:"params,omitempty"`

	// **参数解释**：  应用ID。  **约束限制**：  不涉及。  **取值范围**：  由英文字母和数字组成，且长度为32个字符。  **默认取值**：  不涉及。
	ApplicationId *string `json:"applicationId,omitempty"`
}

func (o RdmParamVoListLifecycleManagedModelUpdateLifecycleStateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoListLifecycleManagedModelUpdateLifecycleStateDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoListLifecycleManagedModelUpdateLifecycleStateDto", string(data)}, " ")
}
