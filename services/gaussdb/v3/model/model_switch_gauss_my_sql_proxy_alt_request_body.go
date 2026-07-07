package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchGaussMySqlProxyAltRequestBody ALT开关对象。
type SwitchGaussMySqlProxyAltRequestBody struct {

	// **参数解释**：  ALT开关。  **约束限制**：  不涉及。  **取值范围**：  - on：开启。  - off：关闭。  **默认取值**：  不涉及。
	AltEnabled string `json:"alt_enabled"`

	// **参数解释**：  只读ALT开关。  **约束限制**：  不涉及。  **取值范围**：  - on：开启。  - off：关闭。  **默认取值**：  off。
	AltForReadonly *string `json:"alt_for_readonly,omitempty"`
}

func (o SwitchGaussMySqlProxyAltRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchGaussMySqlProxyAltRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchGaussMySqlProxyAltRequestBody", string(data)}, " ")
}
