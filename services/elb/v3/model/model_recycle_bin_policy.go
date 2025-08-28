package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecycleBinPolicy **参数解释**：回收站的回收配置。
type RecycleBinPolicy struct {

	// **参数解释**：允许进入回收站的最小创建时间，不足此时长则删除时不满足进入回收站的条件。  **取值范围**：不涉及
	RecycleThresholdHour *int32 `json:"recycle_threshold_hour,omitempty"`

	// **参数解释**：进入回收站的最大保留时长。  **取值范围**：不涉及
	RetentionHour *int32 `json:"retention_hour,omitempty"`
}

func (o RecycleBinPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecycleBinPolicy struct{}"
	}

	return strings.Join([]string{"RecycleBinPolicy", string(data)}, " ")
}
