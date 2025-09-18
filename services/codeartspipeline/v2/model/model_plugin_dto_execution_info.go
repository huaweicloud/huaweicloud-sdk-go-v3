package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PluginDtoExecutionInfo **参数解释**： 执行信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type PluginDtoExecutionInfo struct {
	InnerExecutionInfo *PluginDtoExecutionInfoInnerExecutionInfo `json:"inner_execution_info"`
}

func (o PluginDtoExecutionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginDtoExecutionInfo struct{}"
	}

	return strings.Join([]string{"PluginDtoExecutionInfo", string(data)}, " ")
}
