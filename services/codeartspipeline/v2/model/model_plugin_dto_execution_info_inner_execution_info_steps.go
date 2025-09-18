package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PluginDtoExecutionInfoInnerExecutionInfoSteps struct {

	// **参数解释**： 任务名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 任务类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Task *string `json:"task,omitempty"`

	// **参数解释**： 参数信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Variables map[string]interface{} `json:"variables,omitempty"`
}

func (o PluginDtoExecutionInfoInnerExecutionInfoSteps) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginDtoExecutionInfoInnerExecutionInfoSteps struct{}"
	}

	return strings.Join([]string{"PluginDtoExecutionInfoInnerExecutionInfoSteps", string(data)}, " ")
}
