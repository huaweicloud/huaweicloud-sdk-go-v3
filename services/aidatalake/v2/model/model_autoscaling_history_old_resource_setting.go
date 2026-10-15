package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AutoscalingHistoryOldResourceSetting **参数解释**：变更前的资源配置。 **取值范围**：不涉及。
type AutoscalingHistoryOldResourceSetting struct {

	// **参数解释**：变更前各算子的并行度，key为算子ID，value为并行度。 **取值范围**：不涉及。
	Parallelism map[string]int32 `json:"parallelism,omitempty"`
}

func (o AutoscalingHistoryOldResourceSetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoscalingHistoryOldResourceSetting struct{}"
	}

	return strings.Join([]string{"AutoscalingHistoryOldResourceSetting", string(data)}, " ")
}
