package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LocalAreaIdDef 本端大区。云连接支持的大区有： - Chinese-Mainland（中国大陆） - Asia-Pacific（亚太） - Africa（非洲） - Western-Latin-America（拉美西） - Eastern-Latin-America（拉美东） - Northern-Latin-America（拉美北）
type LocalAreaIdDef struct {
	value string
}

type LocalAreaIdDefEnum struct {
	CHINESE_MAINLAND       LocalAreaIdDef
	ASIA_PACIFIC           LocalAreaIdDef
	AFRICA                 LocalAreaIdDef
	WESTERN_LATIN_AMERICA  LocalAreaIdDef
	EASTERN_LATIN_AMERICA  LocalAreaIdDef
	NORTHERN_LATIN_AMERICA LocalAreaIdDef
}

func GetLocalAreaIdDefEnum() LocalAreaIdDefEnum {
	return LocalAreaIdDefEnum{
		CHINESE_MAINLAND: LocalAreaIdDef{
			value: "Chinese-Mainland",
		},
		ASIA_PACIFIC: LocalAreaIdDef{
			value: "Asia-Pacific",
		},
		AFRICA: LocalAreaIdDef{
			value: "Africa",
		},
		WESTERN_LATIN_AMERICA: LocalAreaIdDef{
			value: "Western-Latin-America",
		},
		EASTERN_LATIN_AMERICA: LocalAreaIdDef{
			value: "Eastern-Latin-America",
		},
		NORTHERN_LATIN_AMERICA: LocalAreaIdDef{
			value: "Northern-Latin-America",
		},
	}
}

func (c LocalAreaIdDef) Value() string {
	return c.value
}

func (c LocalAreaIdDef) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LocalAreaIdDef) UnmarshalJSON(b []byte) error {
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
