package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ImageChipType struct {
	value string
}

type ImageChipTypeEnum struct {
	X86 ImageChipType
	ARM ImageChipType
}

func GetImageChipTypeEnum() ImageChipTypeEnum {
	return ImageChipTypeEnum{
		X86: ImageChipType{
			value: "X86",
		},
		ARM: ImageChipType{
			value: "ARM",
		},
	}
}

func (c ImageChipType) Value() string {
	return c.value
}

func (c ImageChipType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageChipType) UnmarshalJSON(b []byte) error {
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
