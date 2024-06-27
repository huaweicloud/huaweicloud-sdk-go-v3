package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClickHouseSlowLogSensitiveStatusResponse Response Object
type UpdateClickHouseSlowLogSensitiveStatusResponse struct {

	// 慢日志脱敏开关。
	OpenSlowLogSwitch *string `json:"open_slow_log_switch,omitempty"`
	HttpStatusCode    int     `json:"-"`
}

func (o UpdateClickHouseSlowLogSensitiveStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClickHouseSlowLogSensitiveStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateClickHouseSlowLogSensitiveStatusResponse", string(data)}, " ")
}
