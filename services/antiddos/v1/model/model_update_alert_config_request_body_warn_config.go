package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 告警配置信息。
type UpdateAlertConfigRequestBodyWarnConfig struct {

	// DDoS攻击
	AntiDDoS bool `json:"antiDDoS" xml:"antiDDoS"`

	// 网页后门
	BackDoors *bool `json:"back_doors,omitempty" xml:"back_doors"`

	// 暴力破解（系统登录，FTP，DB）
	BruceForce *bool `json:"bruce_force,omitempty" xml:"bruce_force"`

	// 数据库进程权限过高
	HighPrivilege *bool `json:"high_privilege,omitempty" xml:"high_privilege"`

	// 异地登录提醒
	RemoteLogin *bool `json:"remote_login,omitempty" xml:"remote_login"`

	// 取值范围： - 0：表示每天一次 - 1：表示半小时一次  对于HID必选。
	SendFrequency *int32 `json:"send_frequency,omitempty" xml:"send_frequency"`

	// 保留字段
	Waf *bool `json:"waf,omitempty" xml:"waf"`

	// 弱口令（系统，数据库）
	WeakPassword *bool `json:"weak_password,omitempty" xml:"weak_password"`
}

func (o UpdateAlertConfigRequestBodyWarnConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlertConfigRequestBodyWarnConfig struct{}"
	}

	return strings.Join([]string{"UpdateAlertConfigRequestBodyWarnConfig", string(data)}, " ")
}
