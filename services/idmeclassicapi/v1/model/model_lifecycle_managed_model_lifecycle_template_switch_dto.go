package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LifecycleManagedModelLifecycleTemplateSwitchDto struct {

	// **参数解释：**  数据实例ID。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id string `json:"id"`

	LifecycleTemplate *ObjectReferenceParamDto `json:"lifecycleTemplate"`

	LifecycleState *ObjectReferenceParamDto `json:"lifecycleState"`
}

func (o LifecycleManagedModelLifecycleTemplateSwitchDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LifecycleManagedModelLifecycleTemplateSwitchDto struct{}"
	}

	return strings.Join([]string{"LifecycleManagedModelLifecycleTemplateSwitchDto", string(data)}, " ")
}
