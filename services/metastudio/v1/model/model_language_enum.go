package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LanguageEnum 智能交互语言 * zh_CN：简体中文（已下线，请使用CN） * en_US：英语（已下线，请使用EN） * CN：中文 * EN：英文 * ESP：西班牙语（仅海外站点支持） * por：葡萄牙语（仅海外站点支持） * Arabic：阿拉伯语（仅海外站点支持） * Thai：泰语（仅海外站点支持）
type LanguageEnum struct {
	value string
}

type LanguageEnumEnum struct {
	ZH_CN  LanguageEnum
	EN_US  LanguageEnum
	CN     LanguageEnum
	EN     LanguageEnum
	ESP    LanguageEnum
	POR    LanguageEnum
	ARABIC LanguageEnum
	THAI   LanguageEnum
}

func GetLanguageEnumEnum() LanguageEnumEnum {
	return LanguageEnumEnum{
		ZH_CN: LanguageEnum{
			value: "zh_CN",
		},
		EN_US: LanguageEnum{
			value: "en_US",
		},
		CN: LanguageEnum{
			value: "CN",
		},
		EN: LanguageEnum{
			value: "EN",
		},
		ESP: LanguageEnum{
			value: "ESP",
		},
		POR: LanguageEnum{
			value: "por",
		},
		ARABIC: LanguageEnum{
			value: "Arabic",
		},
		THAI: LanguageEnum{
			value: "Thai",
		},
	}
}

func (c LanguageEnum) Value() string {
	return c.value
}

func (c LanguageEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LanguageEnum) UnmarshalJSON(b []byte) error {
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
