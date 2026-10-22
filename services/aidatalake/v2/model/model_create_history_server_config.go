package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHistoryServerConfig **参数解释**：History Server配置。
type CreateHistoryServerConfig struct {

	// **参数解释**：存储历史服务器功能的启用/禁用状态。必须将其设置为“true”才能启用该功能。 **约束限制**：不涉及。 **取值范围**：   - true：启用。   - false：禁用。 **默认取值**：false。
	Enabled *bool `json:"enabled,omitempty"`

	Obs *ObsConfig `json:"obs,omitempty"`
}

func (o CreateHistoryServerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHistoryServerConfig struct{}"
	}

	return strings.Join([]string{"CreateHistoryServerConfig", string(data)}, " ")
}
