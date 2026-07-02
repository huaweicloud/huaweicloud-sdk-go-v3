package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetDdlLogPolicyRequestBody 设置DDL日志下载策略
type SetDdlLogPolicyRequestBody struct {

	// **参数解释**：  设置DDL日志下载功能开关。  **约束限制**：  不涉及。  **取值范围**：  - ON，开启。 - OFF，关闭。  **默认取值**：  不涉及。
	SwitchStatus string `json:"switch_status"`

	// **参数解释**：  设置DDL日志保留天数。  **约束限制**：  不涉及。  **取值范围**：  1~30。  **默认取值**：  3。
	KeepDays int32 `json:"keep_days"`
}

func (o SetDdlLogPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetDdlLogPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"SetDdlLogPolicyRequestBody", string(data)}, " ")
}
