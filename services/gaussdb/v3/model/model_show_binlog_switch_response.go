package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBinlogSwitchResponse Response Object
type ShowBinlogSwitchResponse struct {

	// **参数解释**：  Binlog日志开关状态。   **取值范围**：  - ON：开启。 - OFF：关闭。
	SwitchStatus *string `json:"switch_status,omitempty"`

	// **参数解释**：  Binlog日志保留天数。  **取值范围**：  1-180。如果为null表示未开启Binlog日志开关。
	KeepDays       *int32 `json:"keep_days,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowBinlogSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBinlogSwitchResponse struct{}"
	}

	return strings.Join([]string{"ShowBinlogSwitchResponse", string(data)}, " ")
}
