package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HistoryServerConfig **参数解释**：History Server配置。
type HistoryServerConfig struct {

	// **参数解释**：存储历史服务器功能的启用/禁用状态。必须将其设置为“true”才能启用该功能。 **约束限制**：不涉及。 **取值范围**：   - true：启用。   - false：禁用。 **默认取值**：false。
	Enabled *bool `json:"enabled,omitempty"`

	Obs *ObsConfig `json:"obs,omitempty"`

	// **参数解释**：Flow Insight是一项诊断与可观测性功能，它从Ray集群收集指标、日志及工作流执行数据，并通过History Server收集器将其传输至已配置的OBS存储。该功能仅在“enabled”参数设置为“true”时可用。 **约束限制**：不涉及。 **取值范围**：   - true：启用。   - false：禁用。 **默认取值**：false。
	FlowInsightEnabled *bool `json:"flow_insight_enabled,omitempty"`
}

func (o HistoryServerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HistoryServerConfig struct{}"
	}

	return strings.Join([]string{"HistoryServerConfig", string(data)}, " ")
}
