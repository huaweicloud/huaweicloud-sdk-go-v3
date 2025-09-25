package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStarRocksSlowlogSensitiveStatusResponse Response Object
type ShowStarRocksSlowlogSensitiveStatusResponse struct {

	// **参数解释**:  慢日志开关状态。  **取值范围**：  - true，开启。 - false，关闭。
	OpenSlowLogSwitch *string `json:"open_slow_log_switch,omitempty"`
	HttpStatusCode    int     `json:"-"`
}

func (o ShowStarRocksSlowlogSensitiveStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStarRocksSlowlogSensitiveStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowStarRocksSlowlogSensitiveStatusResponse", string(data)}, " ")
}
