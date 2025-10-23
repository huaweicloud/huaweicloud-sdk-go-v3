package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ProtectedActionSwitcher struct {

	// **参数解释：** 开关名。 **约束限制：** 不涉及。 **取值范围：** - allowed_force_push，允许强制推送（默认不允许强推）。 **默认取值：** 不涉及。
	Name *ProtectedActionSwitcherName `json:"name,omitempty"`

	// **参数解释：** 是否开启。 **约束限制：** 不涉及。 **取值范围：** - true，开启 - false，关闭 **默认取值：** 不涉及。
	Enable *bool `json:"enable,omitempty"`
}

func (o ProtectedActionSwitcher) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectedActionSwitcher struct{}"
	}

	return strings.Join([]string{"ProtectedActionSwitcher", string(data)}, " ")
}

type ProtectedActionSwitcherName struct {
	value string
}

type ProtectedActionSwitcherNameEnum struct {
	ALLOWED_FORCE_PUSH ProtectedActionSwitcherName
}

func GetProtectedActionSwitcherNameEnum() ProtectedActionSwitcherNameEnum {
	return ProtectedActionSwitcherNameEnum{
		ALLOWED_FORCE_PUSH: ProtectedActionSwitcherName{
			value: "allowed_force_push",
		},
	}
}

func (c ProtectedActionSwitcherName) Value() string {
	return c.value
}

func (c ProtectedActionSwitcherName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProtectedActionSwitcherName) UnmarshalJSON(b []byte) error {
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
