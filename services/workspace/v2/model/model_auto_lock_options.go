package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AutoLockOptions struct {

	// 锁屏等待时间（分钟）。取值范围为[3-86400]。默认：10。
	AutoLockMinutes *int32 `json:"auto_lock_minutes,omitempty"`

	// 自动断开或注销。取值为：AUTO_DISCONNECT：自动断开。AUTO_LOGOUT：自动注销。DISABLED：已禁用。（默认）AUTO_RESTART：自动重启。AUTO_STOP：自动停止。HIBERNATE:休眠。
	AutoDisconnect *AutoLockOptionsAutoDisconnect `json:"auto_disconnect,omitempty"`

	Options *AutoDisconnectOrLogoutControlOptions `json:"options,omitempty"`
}

func (o AutoLockOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoLockOptions struct{}"
	}

	return strings.Join([]string{"AutoLockOptions", string(data)}, " ")
}

type AutoLockOptionsAutoDisconnect struct {
	value string
}

type AutoLockOptionsAutoDisconnectEnum struct {
	AUTO_DISCONNECT AutoLockOptionsAutoDisconnect
	AUTO_LOGOUT     AutoLockOptionsAutoDisconnect
	DISABLED        AutoLockOptionsAutoDisconnect
	AUTO_RESTART    AutoLockOptionsAutoDisconnect
	AUTO_STOP       AutoLockOptionsAutoDisconnect
	HIBERNATE       AutoLockOptionsAutoDisconnect
}

func GetAutoLockOptionsAutoDisconnectEnum() AutoLockOptionsAutoDisconnectEnum {
	return AutoLockOptionsAutoDisconnectEnum{
		AUTO_DISCONNECT: AutoLockOptionsAutoDisconnect{
			value: "AUTO_DISCONNECT",
		},
		AUTO_LOGOUT: AutoLockOptionsAutoDisconnect{
			value: "AUTO_LOGOUT",
		},
		DISABLED: AutoLockOptionsAutoDisconnect{
			value: "DISABLED",
		},
		AUTO_RESTART: AutoLockOptionsAutoDisconnect{
			value: "AUTO_RESTART",
		},
		AUTO_STOP: AutoLockOptionsAutoDisconnect{
			value: "AUTO_STOP",
		},
		HIBERNATE: AutoLockOptionsAutoDisconnect{
			value: "HIBERNATE",
		},
	}
}

func (c AutoLockOptionsAutoDisconnect) Value() string {
	return c.value
}

func (c AutoLockOptionsAutoDisconnect) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AutoLockOptionsAutoDisconnect) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
