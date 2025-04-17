package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AsrTypeEnum 对接的ASR厂商类型 * EI_SIS:华为云SIS（仅国内站点支持） * MOBVOI:出门问问（仅海外站点支持）
type AsrTypeEnum struct {
	value string
}

type AsrTypeEnumEnum struct {
	EI_SIS AsrTypeEnum
	MOBVOI AsrTypeEnum
}

func GetAsrTypeEnumEnum() AsrTypeEnumEnum {
	return AsrTypeEnumEnum{
		EI_SIS: AsrTypeEnum{
			value: "EI_SIS",
		},
		MOBVOI: AsrTypeEnum{
			value: "MOBVOI",
		},
	}
}

func (c AsrTypeEnum) Value() string {
	return c.value
}

func (c AsrTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AsrTypeEnum) UnmarshalJSON(b []byte) error {
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
