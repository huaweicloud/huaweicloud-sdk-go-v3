package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SbcAutomaticDisconnectionOptions struct {

	// 等待时间（分钟）
	DisconnectionWaitingTime *int32 `json:"disconnection_waiting_time,omitempty"`

	// 是否自动注销。取值为：false：表示是。true：表示否。
	SbcAutoLogout *bool `json:"sbc_auto_logout,omitempty"`

	AutoLogoutOptions *AutoLogoutOptions `json:"auto_logout_options,omitempty"`
}

func (o SbcAutomaticDisconnectionOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SbcAutomaticDisconnectionOptions struct{}"
	}

	return strings.Join([]string{"SbcAutomaticDisconnectionOptions", string(data)}, " ")
}
