package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AutoStopConfig struct {

	// **参数解释**：自动停止集群开关。 **约束限制**：不涉及。 **取值范围**：   - true：开启自动停止集群。   - false：关闭自动停止集群。 **默认取值**：不涉及。
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**：指定时间不活动后自动停止集群，单位分钟。 **约束限制**：不涉及。 **取值范围**：5~360。 **默认取值**：10。
	IdleTimeoutMinutes *int32 `json:"idle_timeout_minutes,omitempty"`
}

func (o AutoStopConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoStopConfig struct{}"
	}

	return strings.Join([]string{"AutoStopConfig", string(data)}, " ")
}
