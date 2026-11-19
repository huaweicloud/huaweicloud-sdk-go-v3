package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchOnColdTableRequestBody struct {

	// **参数解释**：  冷热分离开关状态。  **约束限制**：  开启冷热数据分离功能后，暂不支持关闭。  **取值范围**：  true：开启。  **默认取值**：  不涉及。
	SwitchOn bool `json:"switch_on"`
}

func (o SwitchOnColdTableRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchOnColdTableRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchOnColdTableRequestBody", string(data)}, " ")
}
