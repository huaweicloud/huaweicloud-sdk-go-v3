package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClickHouseSlowLogSensitiveStatusResponse Response Object
type ShowClickHouseSlowLogSensitiveStatusResponse struct {

	// 慢日志脱敏开关。
	OpenSlowLogSwitch *string `json:"open_slow_log_switch,omitempty"`
	HttpStatusCode    int     `json:"-"`
}

func (o ShowClickHouseSlowLogSensitiveStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClickHouseSlowLogSensitiveStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowClickHouseSlowLogSensitiveStatusResponse", string(data)}, " ")
}
