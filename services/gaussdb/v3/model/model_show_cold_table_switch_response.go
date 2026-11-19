package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowColdTableSwitchResponse Response Object
type ShowColdTableSwitchResponse struct {

	// **参数解释**：  冷热分离开关状态。  **取值范围**： - off：关闭。 - on：开启。
	SwitchStatus *string `json:"switch_status,omitempty"`

	// **参数解释**：  支持冷热分离的最低内核版本。  **取值范围**：  不涉及。
	KernelVersion  *string `json:"kernel_version,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowColdTableSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowColdTableSwitchResponse struct{}"
	}

	return strings.Join([]string{"ShowColdTableSwitchResponse", string(data)}, " ")
}
