package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// V2ImageVersionType **参数解释**：版本类型。 **约束限制**：不涉及。 **取值范围**：   - STABLE：稳定版本。   - RECOMMEND：推荐版本。   - BETA：测试版本。 **默认取值**：不涉及。
type V2ImageVersionType struct {
	value string
}

type V2ImageVersionTypeEnum struct {
	STABLE    V2ImageVersionType
	RECOMMEND V2ImageVersionType
	BETA      V2ImageVersionType
}

func GetV2ImageVersionTypeEnum() V2ImageVersionTypeEnum {
	return V2ImageVersionTypeEnum{
		STABLE: V2ImageVersionType{
			value: "STABLE",
		},
		RECOMMEND: V2ImageVersionType{
			value: "RECOMMEND",
		},
		BETA: V2ImageVersionType{
			value: "BETA",
		},
	}
}

func (c V2ImageVersionType) Value() string {
	return c.value
}

func (c V2ImageVersionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *V2ImageVersionType) UnmarshalJSON(b []byte) error {
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
