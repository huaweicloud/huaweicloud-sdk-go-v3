package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DisconnectLogoutOptions struct {

	// 断开后自动注销等待时间（分钟）。取值范围为[3-86400]。默认：10。
	DisconnectLogoutMinutes *int32 `json:"disconnect_logout_minutes,omitempty"`
}

func (o DisconnectLogoutOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisconnectLogoutOptions struct{}"
	}

	return strings.Join([]string{"DisconnectLogoutOptions", string(data)}, " ")
}
