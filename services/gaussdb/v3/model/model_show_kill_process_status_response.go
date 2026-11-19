package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKillProcessStatusResponse Response Object
type ShowKillProcessStatusResponse struct {

	// **参数解释**：  自动Kill会话开关状态。  **取值范围**：  - ON：开启。 - OFF：关闭。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowKillProcessStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKillProcessStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowKillProcessStatusResponse", string(data)}, " ")
}
