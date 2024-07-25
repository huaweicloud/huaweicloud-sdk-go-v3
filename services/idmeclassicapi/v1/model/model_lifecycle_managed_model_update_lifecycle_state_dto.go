package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LifecycleManagedModelUpdateLifecycleStateDto struct {

	// **参数解释：**  数据实例ID。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id string `json:"id"`

	LifecycleState *ObjectReferenceParamDto `json:"lifecycleState"`

	// **参数解释：**  修改人。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`
}

func (o LifecycleManagedModelUpdateLifecycleStateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LifecycleManagedModelUpdateLifecycleStateDto struct{}"
	}

	return strings.Join([]string{"LifecycleManagedModelUpdateLifecycleStateDto", string(data)}, " ")
}
