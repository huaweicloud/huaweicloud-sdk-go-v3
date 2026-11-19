package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAutoKillProcessStatusRequestBody struct {

	// **参数解释**：  开关状态。  **约束限制**：  不涉及。  **取值范围**：  - ON：开启。 - OFF：关闭。  **默认取值**：  不涉及。
	Status string `json:"status"`
}

func (o UpdateAutoKillProcessStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAutoKillProcessStatusRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAutoKillProcessStatusRequestBody", string(data)}, " ")
}
