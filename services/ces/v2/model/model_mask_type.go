package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MaskType 屏蔽类型。START_END_TIME：按起止时间屏蔽，FOREVER_TIME：永久时间屏蔽，CYCLE_TIME：按周期时间屏蔽。
type MaskType struct {
	value string
}

type MaskTypeEnum struct {
	START_END_TIME MaskType
	FOREVER_TIME   MaskType
	CYCLE_TIME     MaskType
}

func GetMaskTypeEnum() MaskTypeEnum {
	return MaskTypeEnum{
		START_END_TIME: MaskType{
			value: "START_END_TIME",
		},
		FOREVER_TIME: MaskType{
			value: "FOREVER_TIME",
		},
		CYCLE_TIME: MaskType{
			value: "CYCLE_TIME",
		},
	}
}

func (c MaskType) Value() string {
	return c.value
}

func (c MaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MaskType) UnmarshalJSON(b []byte) error {
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
