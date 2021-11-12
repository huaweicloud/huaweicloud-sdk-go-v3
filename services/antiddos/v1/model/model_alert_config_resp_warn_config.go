package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 告警配置信息
type AlertConfigRespWarnConfig struct {
	// DDoS攻击

	AntiDDoS bool `json:"antiDDoS"`
	// 网页后门

	BackDoors *bool `json:"back_doors,omitempty"`
	// 暴力破解（系统登录，FTP，DB）

	BruceForce *bool `json:"bruce_force,omitempty"`
	// 数据库进程权限过高

	HighPrivilege *bool `json:"high_privilege,omitempty"`
	// 异地登录提醒

	RemoteLogin *bool `json:"remote_login,omitempty"`
	// 取值范围： - 0：表示每天一次 - 1：表示半小时一次  对于HID必选。

	SendFrequency *int32 `json:"send_frequency,omitempty"`
	// 保留字段

	Waf *bool `json:"waf,omitempty"`
	// 弱口令（系统，数据库）

	WeakPassword *bool `json:"weak_password,omitempty"`
}

func (o AlertConfigRespWarnConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertConfigRespWarnConfig struct{}"
	}

	return strings.Join([]string{"AlertConfigRespWarnConfig", string(data)}, " ")
}
