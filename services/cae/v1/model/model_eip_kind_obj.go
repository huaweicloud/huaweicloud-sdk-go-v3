package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EipKindObj API类型，固定值“eip”，该值不可修改。
type EipKindObj struct {
	value string
}

type EipKindObjEnum struct {
	EIP EipKindObj
}

func GetEipKindObjEnum() EipKindObjEnum {
	return EipKindObjEnum{
		EIP: EipKindObj{
			value: "eip",
		},
	}
}

func (c EipKindObj) Value() string {
	return c.value
}

func (c EipKindObj) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EipKindObj) UnmarshalJSON(b []byte) error {
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
