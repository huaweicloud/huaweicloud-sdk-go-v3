package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSlowlogDesensitizationSwitchResponse Response Object
type ShowSlowlogDesensitizationSwitchResponse struct {

	// 开启或关闭慢日志脱敏，取值为on或off。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSlowlogDesensitizationSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowlogDesensitizationSwitchResponse struct{}"
	}

	return strings.Join([]string{"ShowSlowlogDesensitizationSwitchResponse", string(data)}, " ")
}
