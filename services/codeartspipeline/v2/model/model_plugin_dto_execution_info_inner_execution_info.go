package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PluginDtoExecutionInfoInnerExecutionInfo **参数解释**： 插件执行信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type PluginDtoExecutionInfoInnerExecutionInfo struct {

	// **参数解释**： 插件类型。 **约束限制**： 不涉及。 **取值范围**： CONTAINER, ZIP, SHELL, COMPOSITE。 **默认取值**： 不涉及。
	ExecutionType *string `json:"execution_type,omitempty"`

	Steps *[]PluginDtoExecutionInfoInnerExecutionInfoSteps `json:"steps,omitempty"`
}

func (o PluginDtoExecutionInfoInnerExecutionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginDtoExecutionInfoInnerExecutionInfo struct{}"
	}

	return strings.Join([]string{"PluginDtoExecutionInfoInnerExecutionInfo", string(data)}, " ")
}
