package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSlowlogSensitiveStatusRequestBody struct {

	// **参数解释**：  慢日志开关状态。  **约束限制**：  不涉及。  **取值范围**： - true，开启。 - false，关闭。            **默认取值**：  不涉及。
	OpenSlowLogSwitch bool `json:"open_slow_log_switch"`
}

func (o UpdateSlowlogSensitiveStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSlowlogSensitiveStatusRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSlowlogSensitiveStatusRequestBody", string(data)}, " ")
}
