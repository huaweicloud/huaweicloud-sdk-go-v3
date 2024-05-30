package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ApplyBgEnum 适用bg。
type ApplyBgEnum struct {
	value string
}

type ApplyBgEnumEnum struct {
	CLOUD_BU ApplyBgEnum
	BG1      ApplyBgEnum
	BG2      ApplyBgEnum
	BG3      ApplyBgEnum
}

func GetApplyBgEnumEnum() ApplyBgEnumEnum {
	return ApplyBgEnumEnum{
		CLOUD_BU: ApplyBgEnum{
			value: "CLOUD_BU",
		},
		BG1: ApplyBgEnum{
			value: "BG1",
		},
		BG2: ApplyBgEnum{
			value: "BG2",
		},
		BG3: ApplyBgEnum{
			value: "BG3",
		},
	}
}

func (c ApplyBgEnum) Value() string {
	return c.value
}

func (c ApplyBgEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApplyBgEnum) UnmarshalJSON(b []byte) error {
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
