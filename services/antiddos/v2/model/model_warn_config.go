package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// WarnConfig 告警配置信息。
type WarnConfig struct {

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
	SendFrequency *WarnConfigSendFrequency `json:"send_frequency,omitempty"`

	// 保留字段
	Waf *bool `json:"waf,omitempty"`

	// 弱口令（系统，数据库）
	WeakPassword *bool `json:"weak_password,omitempty"`
}

func (o WarnConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WarnConfig struct{}"
	}

	return strings.Join([]string{"WarnConfig", string(data)}, " ")
}

type WarnConfigSendFrequency struct {
	value int32
}

type WarnConfigSendFrequencyEnum struct {
	E_0 WarnConfigSendFrequency
	E_1 WarnConfigSendFrequency
}

func GetWarnConfigSendFrequencyEnum() WarnConfigSendFrequencyEnum {
	return WarnConfigSendFrequencyEnum{
		E_0: WarnConfigSendFrequency{
			value: 0,
		}, E_1: WarnConfigSendFrequency{
			value: 1,
		},
	}
}

func (c WarnConfigSendFrequency) Value() int32 {
	return c.value
}

func (c WarnConfigSendFrequency) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WarnConfigSendFrequency) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
