package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTerminalsBindingDesktopsConfigResponse Response Object
type ListTerminalsBindingDesktopsConfigResponse struct {

	// 绑定开关，只取值ON或OFF。
	TcBindSwitch *ListTerminalsBindingDesktopsConfigResponseTcBindSwitch `json:"tc_bind_switch,omitempty"`

	// 自动绑定开关，只取值ON或OFF。
	TcAutoBindSwitch *ListTerminalsBindingDesktopsConfigResponseTcAutoBindSwitch `json:"tc_auto_bind_switch,omitempty"`

	// 最大绑定数量，默认值为1。
	TcAutoBindMax  *int32 `json:"tc_auto_bind_max,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTerminalsBindingDesktopsConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTerminalsBindingDesktopsConfigResponse struct{}"
	}

	return strings.Join([]string{"ListTerminalsBindingDesktopsConfigResponse", string(data)}, " ")
}

type ListTerminalsBindingDesktopsConfigResponseTcBindSwitch struct {
	value string
}

type ListTerminalsBindingDesktopsConfigResponseTcBindSwitchEnum struct {
	ON  ListTerminalsBindingDesktopsConfigResponseTcBindSwitch
	OFF ListTerminalsBindingDesktopsConfigResponseTcBindSwitch
}

func GetListTerminalsBindingDesktopsConfigResponseTcBindSwitchEnum() ListTerminalsBindingDesktopsConfigResponseTcBindSwitchEnum {
	return ListTerminalsBindingDesktopsConfigResponseTcBindSwitchEnum{
		ON: ListTerminalsBindingDesktopsConfigResponseTcBindSwitch{
			value: "ON",
		},
		OFF: ListTerminalsBindingDesktopsConfigResponseTcBindSwitch{
			value: "OFF",
		},
	}
}

func (c ListTerminalsBindingDesktopsConfigResponseTcBindSwitch) Value() string {
	return c.value
}

func (c ListTerminalsBindingDesktopsConfigResponseTcBindSwitch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTerminalsBindingDesktopsConfigResponseTcBindSwitch) UnmarshalJSON(b []byte) error {
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

type ListTerminalsBindingDesktopsConfigResponseTcAutoBindSwitch struct {
	value string
}

type ListTerminalsBindingDesktopsConfigResponseTcAutoBindSwitchEnum struct {
	ON  ListTerminalsBindingDesktopsConfigResponseTcAutoBindSwitch
	OFF ListTerminalsBindingDesktopsConfigResponseTcAutoBindSwitch
}

func GetListTerminalsBindingDesktopsConfigResponseTcAutoBindSwitchEnum() ListTerminalsBindingDesktopsConfigResponseTcAutoBindSwitchEnum {
	return ListTerminalsBindingDesktopsConfigResponseTcAutoBindSwitchEnum{
		ON: ListTerminalsBindingDesktopsConfigResponseTcAutoBindSwitch{
			value: "ON",
		},
		OFF: ListTerminalsBindingDesktopsConfigResponseTcAutoBindSwitch{
			value: "OFF",
		},
	}
}

func (c ListTerminalsBindingDesktopsConfigResponseTcAutoBindSwitch) Value() string {
	return c.value
}

func (c ListTerminalsBindingDesktopsConfigResponseTcAutoBindSwitch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTerminalsBindingDesktopsConfigResponseTcAutoBindSwitch) UnmarshalJSON(b []byte) error {
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
