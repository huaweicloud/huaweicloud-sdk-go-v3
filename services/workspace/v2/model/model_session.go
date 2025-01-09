package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Session struct {
	Vdi *Vdi `json:"vdi,omitempty"`

	// 是否开启自助维护台抢占登陆。取值为：false：表示关闭。true：表示开启。
	SelfHelpConsole *bool `json:"self_help_console,omitempty"`

	// 是否锁屏后断开
	DisconnectOnLockFlag *bool `json:"disconnect_on_lock_flag,omitempty"`
}

func (o Session) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Session struct{}"
	}

	return strings.Join([]string{"Session", string(data)}, " ")
}
