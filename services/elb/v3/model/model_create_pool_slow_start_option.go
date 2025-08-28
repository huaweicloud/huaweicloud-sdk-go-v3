package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePoolSlowStartOption struct {

	// **参数解释**：慢启动的开关。  **约束限制**：不涉及  **取值范围**： - true：开启。 - false：关闭。  **默认取值**：false
	Enable *bool `json:"enable,omitempty"`

	// **参数解释**：慢启动的持续时间。  **约束限制**：不涉及  **取值范围**：30-1200，单位：秒。  **默认取值**：30
	Duration *int32 `json:"duration,omitempty"`
}

func (o CreatePoolSlowStartOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePoolSlowStartOption struct{}"
	}

	return strings.Join([]string{"CreatePoolSlowStartOption", string(data)}, " ")
}
