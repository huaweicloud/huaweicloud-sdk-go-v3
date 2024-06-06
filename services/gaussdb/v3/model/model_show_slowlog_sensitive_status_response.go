package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSlowlogSensitiveStatusResponse Response Object
type ShowSlowlogSensitiveStatusResponse struct {

	// 慢日志开关状态。
	OpenSlowLogSwitch *bool `json:"open_slow_log_switch,omitempty"`
	HttpStatusCode    int   `json:"-"`
}

func (o ShowSlowlogSensitiveStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowlogSensitiveStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowSlowlogSensitiveStatusResponse", string(data)}, " ")
}
