package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SupportedServiceEnum struct {
	value string
}

type SupportedServiceEnumEnum struct {
	VIDEO_2_D SupportedServiceEnum
	LIVE_2_D  SupportedServiceEnum
	CHAT_2_D  SupportedServiceEnum
}

func GetSupportedServiceEnumEnum() SupportedServiceEnumEnum {
	return SupportedServiceEnumEnum{
		VIDEO_2_D: SupportedServiceEnum{
			value: "VIDEO_2D",
		},
		LIVE_2_D: SupportedServiceEnum{
			value: "LIVE_2D",
		},
		CHAT_2_D: SupportedServiceEnum{
			value: "CHAT_2D",
		},
	}
}

func (c SupportedServiceEnum) Value() string {
	return c.value
}

func (c SupportedServiceEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SupportedServiceEnum) UnmarshalJSON(b []byte) error {
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
