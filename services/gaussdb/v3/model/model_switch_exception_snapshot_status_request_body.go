package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchExceptionSnapshotStatusRequestBody struct {

	// **参数解释**：  异常快照开关状态。  **约束限制**：  不涉及。  **取值范围**：  - true：开启。 - false：关闭。  **默认取值**：  不涉及。
	Enabled bool `json:"enabled"`
}

func (o SwitchExceptionSnapshotStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchExceptionSnapshotStatusRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchExceptionSnapshotStatusRequestBody", string(data)}, " ")
}
