package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TerminalsBindingDesktopsConfig struct {

	// 绑定开关，只取值ON或OFF。
	TcBindSwitch TerminalsBindingDesktopsConfigTcBindSwitch `json:"tc_bind_switch"`

	// 自动绑定开关，只取值ON或OFF。
	TcAutoBindSwitch *TerminalsBindingDesktopsConfigTcAutoBindSwitch `json:"tc_auto_bind_switch,omitempty"`

	// 最大绑定数量，默认值为1。
	TcAutoBindMax *int32 `json:"tc_auto_bind_max,omitempty"`
}

func (o TerminalsBindingDesktopsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TerminalsBindingDesktopsConfig struct{}"
	}

	return strings.Join([]string{"TerminalsBindingDesktopsConfig", string(data)}, " ")
}

type TerminalsBindingDesktopsConfigTcBindSwitch struct {
	value string
}

type TerminalsBindingDesktopsConfigTcBindSwitchEnum struct {
	ON  TerminalsBindingDesktopsConfigTcBindSwitch
	OFF TerminalsBindingDesktopsConfigTcBindSwitch
}

func GetTerminalsBindingDesktopsConfigTcBindSwitchEnum() TerminalsBindingDesktopsConfigTcBindSwitchEnum {
	return TerminalsBindingDesktopsConfigTcBindSwitchEnum{
		ON: TerminalsBindingDesktopsConfigTcBindSwitch{
			value: "ON",
		},
		OFF: TerminalsBindingDesktopsConfigTcBindSwitch{
			value: "OFF",
		},
	}
}

func (c TerminalsBindingDesktopsConfigTcBindSwitch) Value() string {
	return c.value
}

func (c TerminalsBindingDesktopsConfigTcBindSwitch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TerminalsBindingDesktopsConfigTcBindSwitch) UnmarshalJSON(b []byte) error {
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

type TerminalsBindingDesktopsConfigTcAutoBindSwitch struct {
	value string
}

type TerminalsBindingDesktopsConfigTcAutoBindSwitchEnum struct {
	ON  TerminalsBindingDesktopsConfigTcAutoBindSwitch
	OFF TerminalsBindingDesktopsConfigTcAutoBindSwitch
}

func GetTerminalsBindingDesktopsConfigTcAutoBindSwitchEnum() TerminalsBindingDesktopsConfigTcAutoBindSwitchEnum {
	return TerminalsBindingDesktopsConfigTcAutoBindSwitchEnum{
		ON: TerminalsBindingDesktopsConfigTcAutoBindSwitch{
			value: "ON",
		},
		OFF: TerminalsBindingDesktopsConfigTcAutoBindSwitch{
			value: "OFF",
		},
	}
}

func (c TerminalsBindingDesktopsConfigTcAutoBindSwitch) Value() string {
	return c.value
}

func (c TerminalsBindingDesktopsConfigTcAutoBindSwitch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TerminalsBindingDesktopsConfigTcAutoBindSwitch) UnmarshalJSON(b []byte) error {
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
