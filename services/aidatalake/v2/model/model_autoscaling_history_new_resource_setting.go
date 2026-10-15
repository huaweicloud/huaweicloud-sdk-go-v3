package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AutoscalingHistoryNewResourceSetting **参数解释**：变更后的资源配置。 **取值范围**：不涉及。
type AutoscalingHistoryNewResourceSetting struct {

	// **参数解释**：变更后各算子的并行度，key为算子ID，value为并行度。 **取值范围**：不涉及。
	Parallelism map[string]int32 `json:"parallelism,omitempty"`
}

func (o AutoscalingHistoryNewResourceSetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoscalingHistoryNewResourceSetting struct{}"
	}

	return strings.Join([]string{"AutoscalingHistoryNewResourceSetting", string(data)}, " ")
}
