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
