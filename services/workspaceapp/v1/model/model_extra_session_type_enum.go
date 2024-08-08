package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExtraSessionTypeEnum 付费会话类型 * `GPU` - GPU规格会话 * `CPU` - 普通CPU规格会话
type ExtraSessionTypeEnum struct {
	value string
}

type ExtraSessionTypeEnumEnum struct {
	GPU ExtraSessionTypeEnum
	CPU ExtraSessionTypeEnum
}

func GetExtraSessionTypeEnumEnum() ExtraSessionTypeEnumEnum {
	return ExtraSessionTypeEnumEnum{
		GPU: ExtraSessionTypeEnum{
			value: "GPU",
		},
		CPU: ExtraSessionTypeEnum{
			value: "CPU",
		},
	}
}

func (c ExtraSessionTypeEnum) Value() string {
	return c.value
}

func (c ExtraSessionTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExtraSessionTypeEnum) UnmarshalJSON(b []byte) error {
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
