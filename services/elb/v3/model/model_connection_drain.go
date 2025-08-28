package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConnectionDrain struct {

	// **参数解释**：延迟注销功能开关。  **约束限制**：不涉及  **取值范围**：true 开启，false 关闭。  **默认取值**：不涉及
	Enable *bool `json:"enable,omitempty"`

	// **参数解释**：延迟注销时间。  **约束限制**：不涉及  **取值范围**：10~4000，单位：秒。  **默认取值**：不涉及
	Timeout *int32 `json:"timeout,omitempty"`
}

func (o ConnectionDrain) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionDrain struct{}"
	}

	return strings.Join([]string{"ConnectionDrain", string(data)}, " ")
}
