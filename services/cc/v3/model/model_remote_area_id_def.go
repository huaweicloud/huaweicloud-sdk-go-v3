package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RemoteAreaIdDef 对端大区。云连接支持的大区有： - Chinese-Mainland（中国大陆） - Asia-Pacific（亚太） - Africa（非洲） - Western-Latin-America（拉美西） - Eastern-Latin-America（拉美东） - Northern-Latin-America（拉美北）
type RemoteAreaIdDef struct {
	value string
}

type RemoteAreaIdDefEnum struct {
	CHINESE_MAINLAND       RemoteAreaIdDef
	ASIA_PACIFIC           RemoteAreaIdDef
	AFRICA                 RemoteAreaIdDef
	WESTERN_LATIN_AMERICA  RemoteAreaIdDef
	EASTERN_LATIN_AMERICA  RemoteAreaIdDef
	NORTHERN_LATIN_AMERICA RemoteAreaIdDef
}

func GetRemoteAreaIdDefEnum() RemoteAreaIdDefEnum {
	return RemoteAreaIdDefEnum{
		CHINESE_MAINLAND: RemoteAreaIdDef{
			value: "Chinese-Mainland",
		},
		ASIA_PACIFIC: RemoteAreaIdDef{
			value: "Asia-Pacific",
		},
		AFRICA: RemoteAreaIdDef{
			value: "Africa",
		},
		WESTERN_LATIN_AMERICA: RemoteAreaIdDef{
			value: "Western-Latin-America",
		},
		EASTERN_LATIN_AMERICA: RemoteAreaIdDef{
			value: "Eastern-Latin-America",
		},
		NORTHERN_LATIN_AMERICA: RemoteAreaIdDef{
			value: "Northern-Latin-America",
		},
	}
}

func (c RemoteAreaIdDef) Value() string {
	return c.value
}

func (c RemoteAreaIdDef) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RemoteAreaIdDef) UnmarshalJSON(b []byte) error {
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
