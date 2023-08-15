package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ReceiveModeEnum 验证码接收模式 VMFA：虚拟MFA设备 HMFA：硬件MFA设备
type ReceiveModeEnum struct {
	value string
}

type ReceiveModeEnumEnum struct {
	VMFA ReceiveModeEnum
	HMFA ReceiveModeEnum
}

func GetReceiveModeEnumEnum() ReceiveModeEnumEnum {
	return ReceiveModeEnumEnum{
		VMFA: ReceiveModeEnum{
			value: "VMFA",
		},
		HMFA: ReceiveModeEnum{
			value: "HMFA",
		},
	}
}

func (c ReceiveModeEnum) Value() string {
	return c.value
}

func (c ReceiveModeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReceiveModeEnum) UnmarshalJSON(b []byte) error {
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
