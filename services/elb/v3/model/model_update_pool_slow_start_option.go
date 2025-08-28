package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePoolSlowStartOption struct {

	// **参数解释**：慢启动的开关。  **约束限制**：不涉及  **取值范围**：true 开启，false 关闭。  **默认取值**：不涉及
	Enable *bool `json:"enable,omitempty"`

	// **参数解释**：慢启动的持续时间。  **约束限制**：不涉及  **取值范围**：30~1200，单位：秒。  **默认取值**：不涉及
	Duration *int32 `json:"duration,omitempty"`
}

func (o UpdatePoolSlowStartOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePoolSlowStartOption struct{}"
	}

	return strings.Join([]string{"UpdatePoolSlowStartOption", string(data)}, " ")
}
