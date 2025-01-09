package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AutoDisconnectOrLogoutControlOptions struct {

	// 断开或注销等待时间（分钟）。取值范围为[1-86400]。默认：1440。
	AutoDisconnectMinutes *int32 `json:"auto_disconnect_minutes,omitempty"`
}

func (o AutoDisconnectOrLogoutControlOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoDisconnectOrLogoutControlOptions struct{}"
	}

	return strings.Join([]string{"AutoDisconnectOrLogoutControlOptions", string(data)}, " ")
}
