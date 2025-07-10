package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PoliciesKeyboardMouse 键盘鼠标。
type PoliciesKeyboardMouse struct {

	// 虚拟机鼠标回馈。取值为： SELFADAPTION：自适应鼠标回馈。 FORCE：强制鼠标回馈。 CLOSE：关闭鼠标回馈。
	MouseFeedback *PoliciesKeyboardMouseMouseFeedback `json:"mouse_feedback,omitempty"`

	// 虚拟机鼠标模拟方式。取值为： ABSOLUTE_POSITION：绝对位置模拟。 RELATIVE_POSITION：相对位置模拟。
	MouseSimulationMode *PoliciesKeyboardMouseMouseSimulationMode `json:"mouse_simulation_mode,omitempty"`

	// 虚拟机外部光标反馈。取值为： false：表示关闭。 true：表示开启。
	ExternalCursorFeedback *bool `json:"external_cursor_feedback,omitempty"`

	// 自助维护台抢占登录。取值为： false：表示关闭。 true：表示开启。
	SelfhelpConsoleEnable *bool `json:"selfhelp_console_enable,omitempty"`

	// 客户端鼠标发送间隔。取值范围为[1-30]。默认：30。
	ClientMouseSendInterval *int32 `json:"client_mouse_send_interval,omitempty"`

	// Windows客户端键盘模式。取值为： GLOBAL：Windows客户端键盘全局模式。（默认） WINDOW：Windows客户端键盘窗口模式。
	WindowsClientKeyboardMode *PoliciesKeyboardMouseWindowsClientKeyboardMode `json:"windows_client_keyboard_mode,omitempty"`

	// Windows客户端鼠标模式。取值为： GLOBAL：Windows客户端鼠标全局模式。 WINDOW：Windows客户端鼠标窗口模式。（默认）
	WindowsClientMouseMode *PoliciesKeyboardMouseWindowsClientMouseMode `json:"windows_client_mouse_mode,omitempty"`

	// Linux客户端键盘模式。取值为： EVENT：Linux客户端键盘事件模式。（默认） GRAPH：Linux客户端键盘图形模式。
	LinuxClientKeyboardMode *PoliciesKeyboardMouseLinuxClientKeyboardMode `json:"linux_client_keyboard_mode,omitempty"`

	// Linux客户端鼠标模式。取值为： EVENT：Linux客户端鼠标事件模式。（默认） GRAPH：Linux客户端鼠标窗口模式。
	LinuxClientMouseMode *PoliciesKeyboardMouseLinuxClientMouseMode `json:"linux_client_mouse_mode,omitempty"`

	// 特殊键盘。取值为： false：表示关闭。 true：表示开启。
	SpecialKeyboard *bool `json:"special_keyboard,omitempty"`

	// 游戏操纵杆。取值为： false：表示关闭。 true：表示开启。
	JoyStickFlag *bool `json:"joy_stick_flag,omitempty"`
}

func (o PoliciesKeyboardMouse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesKeyboardMouse struct{}"
	}

	return strings.Join([]string{"PoliciesKeyboardMouse", string(data)}, " ")
}

type PoliciesKeyboardMouseMouseFeedback struct {
	value string
}

type PoliciesKeyboardMouseMouseFeedbackEnum struct {
	SELFADAPTION PoliciesKeyboardMouseMouseFeedback
	FORCE        PoliciesKeyboardMouseMouseFeedback
	CLOSE        PoliciesKeyboardMouseMouseFeedback
}

func GetPoliciesKeyboardMouseMouseFeedbackEnum() PoliciesKeyboardMouseMouseFeedbackEnum {
	return PoliciesKeyboardMouseMouseFeedbackEnum{
		SELFADAPTION: PoliciesKeyboardMouseMouseFeedback{
			value: "SELFADAPTION",
		},
		FORCE: PoliciesKeyboardMouseMouseFeedback{
			value: "FORCE",
		},
		CLOSE: PoliciesKeyboardMouseMouseFeedback{
			value: "CLOSE",
		},
	}
}

func (c PoliciesKeyboardMouseMouseFeedback) Value() string {
	return c.value
}

func (c PoliciesKeyboardMouseMouseFeedback) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesKeyboardMouseMouseFeedback) UnmarshalJSON(b []byte) error {
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

type PoliciesKeyboardMouseMouseSimulationMode struct {
	value string
}

type PoliciesKeyboardMouseMouseSimulationModeEnum struct {
	ABSOLUTE_POSITION PoliciesKeyboardMouseMouseSimulationMode
	RELATIVE_POSITION PoliciesKeyboardMouseMouseSimulationMode
}

func GetPoliciesKeyboardMouseMouseSimulationModeEnum() PoliciesKeyboardMouseMouseSimulationModeEnum {
	return PoliciesKeyboardMouseMouseSimulationModeEnum{
		ABSOLUTE_POSITION: PoliciesKeyboardMouseMouseSimulationMode{
			value: "ABSOLUTE_POSITION",
		},
		RELATIVE_POSITION: PoliciesKeyboardMouseMouseSimulationMode{
			value: "RELATIVE_POSITION",
		},
	}
}

func (c PoliciesKeyboardMouseMouseSimulationMode) Value() string {
	return c.value
}

func (c PoliciesKeyboardMouseMouseSimulationMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesKeyboardMouseMouseSimulationMode) UnmarshalJSON(b []byte) error {
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

type PoliciesKeyboardMouseWindowsClientKeyboardMode struct {
	value string
}

type PoliciesKeyboardMouseWindowsClientKeyboardModeEnum struct {
	GLOBAL PoliciesKeyboardMouseWindowsClientKeyboardMode
	WINDOW PoliciesKeyboardMouseWindowsClientKeyboardMode
}

func GetPoliciesKeyboardMouseWindowsClientKeyboardModeEnum() PoliciesKeyboardMouseWindowsClientKeyboardModeEnum {
	return PoliciesKeyboardMouseWindowsClientKeyboardModeEnum{
		GLOBAL: PoliciesKeyboardMouseWindowsClientKeyboardMode{
			value: "GLOBAL",
		},
		WINDOW: PoliciesKeyboardMouseWindowsClientKeyboardMode{
			value: "WINDOW",
		},
	}
}

func (c PoliciesKeyboardMouseWindowsClientKeyboardMode) Value() string {
	return c.value
}

func (c PoliciesKeyboardMouseWindowsClientKeyboardMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesKeyboardMouseWindowsClientKeyboardMode) UnmarshalJSON(b []byte) error {
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

type PoliciesKeyboardMouseWindowsClientMouseMode struct {
	value string
}

type PoliciesKeyboardMouseWindowsClientMouseModeEnum struct {
	GLOBAL PoliciesKeyboardMouseWindowsClientMouseMode
	WINDOW PoliciesKeyboardMouseWindowsClientMouseMode
}

func GetPoliciesKeyboardMouseWindowsClientMouseModeEnum() PoliciesKeyboardMouseWindowsClientMouseModeEnum {
	return PoliciesKeyboardMouseWindowsClientMouseModeEnum{
		GLOBAL: PoliciesKeyboardMouseWindowsClientMouseMode{
			value: "GLOBAL",
		},
		WINDOW: PoliciesKeyboardMouseWindowsClientMouseMode{
			value: "WINDOW",
		},
	}
}

func (c PoliciesKeyboardMouseWindowsClientMouseMode) Value() string {
	return c.value
}

func (c PoliciesKeyboardMouseWindowsClientMouseMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesKeyboardMouseWindowsClientMouseMode) UnmarshalJSON(b []byte) error {
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

type PoliciesKeyboardMouseLinuxClientKeyboardMode struct {
	value string
}

type PoliciesKeyboardMouseLinuxClientKeyboardModeEnum struct {
	EVENT PoliciesKeyboardMouseLinuxClientKeyboardMode
	GRAPH PoliciesKeyboardMouseLinuxClientKeyboardMode
}

func GetPoliciesKeyboardMouseLinuxClientKeyboardModeEnum() PoliciesKeyboardMouseLinuxClientKeyboardModeEnum {
	return PoliciesKeyboardMouseLinuxClientKeyboardModeEnum{
		EVENT: PoliciesKeyboardMouseLinuxClientKeyboardMode{
			value: "EVENT",
		},
		GRAPH: PoliciesKeyboardMouseLinuxClientKeyboardMode{
			value: "GRAPH",
		},
	}
}

func (c PoliciesKeyboardMouseLinuxClientKeyboardMode) Value() string {
	return c.value
}

func (c PoliciesKeyboardMouseLinuxClientKeyboardMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesKeyboardMouseLinuxClientKeyboardMode) UnmarshalJSON(b []byte) error {
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

type PoliciesKeyboardMouseLinuxClientMouseMode struct {
	value string
}

type PoliciesKeyboardMouseLinuxClientMouseModeEnum struct {
	EVENT PoliciesKeyboardMouseLinuxClientMouseMode
	GRAPH PoliciesKeyboardMouseLinuxClientMouseMode
}

func GetPoliciesKeyboardMouseLinuxClientMouseModeEnum() PoliciesKeyboardMouseLinuxClientMouseModeEnum {
	return PoliciesKeyboardMouseLinuxClientMouseModeEnum{
		EVENT: PoliciesKeyboardMouseLinuxClientMouseMode{
			value: "EVENT",
		},
		GRAPH: PoliciesKeyboardMouseLinuxClientMouseMode{
			value: "GRAPH",
		},
	}
}

func (c PoliciesKeyboardMouseLinuxClientMouseMode) Value() string {
	return c.value
}

func (c PoliciesKeyboardMouseLinuxClientMouseMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesKeyboardMouseLinuxClientMouseMode) UnmarshalJSON(b []byte) error {
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
