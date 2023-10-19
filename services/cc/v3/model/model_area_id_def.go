package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AreaIdDef 本端大区。 云连接支持的大区有： | Chinese-Mainland | 中国大陆 | | Asia-Pacific | 亚太 | | Africa | 非洲 | | Western-Latin-America | 拉美西 | | Eastern-Latin-America | 拉美东 | | Northern-Latin-America | 拉美北 |
type AreaIdDef struct {
	value string
}

type AreaIdDefEnum struct {
	CHINESE_MAINLAND       AreaIdDef
	ASIA_PACIFIC           AreaIdDef
	AFRICA                 AreaIdDef
	WESTERN_LATIN_AMERICA  AreaIdDef
	EASTERN_LATIN_AMERICA  AreaIdDef
	NORTHERN_LATIN_AMERICA AreaIdDef
}

func GetAreaIdDefEnum() AreaIdDefEnum {
	return AreaIdDefEnum{
		CHINESE_MAINLAND: AreaIdDef{
			value: "Chinese-Mainland",
		},
		ASIA_PACIFIC: AreaIdDef{
			value: "Asia-Pacific",
		},
		AFRICA: AreaIdDef{
			value: "Africa",
		},
		WESTERN_LATIN_AMERICA: AreaIdDef{
			value: "Western-Latin-America",
		},
		EASTERN_LATIN_AMERICA: AreaIdDef{
			value: "Eastern-Latin-America",
		},
		NORTHERN_LATIN_AMERICA: AreaIdDef{
			value: "Northern-Latin-America",
		},
	}
}

func (c AreaIdDef) Value() string {
	return c.value
}

func (c AreaIdDef) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AreaIdDef) UnmarshalJSON(b []byte) error {
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
