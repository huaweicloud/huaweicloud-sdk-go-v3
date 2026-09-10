package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Platform **参数解释**： 部署平台。 **约束限制**： 不涉及 **取值范围**： * Modelarts：ModelArts平台 * CCE：CCE平台 **默认取值**： 不涉及
type Platform struct {
	value string
}

type PlatformEnum struct {
	MODELARTS Platform
	CCE       Platform
}

func GetPlatformEnum() PlatformEnum {
	return PlatformEnum{
		MODELARTS: Platform{
			value: "Modelarts",
		},
		CCE: Platform{
			value: "CCE",
		},
	}
}

func (c Platform) Value() string {
	return c.value
}

func (c Platform) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Platform) UnmarshalJSON(b []byte) error {
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
