package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResourceType **参数解释**：资源使用模式。 **约束限制**：不涉及。 **取值范围**：   - RESERVED：预留资源。   - ONDEMAND：按需弹性。   - HYBRID：混合模式。 **默认取值**：不涉及。
type ResourceType struct {
	value string
}

type ResourceTypeEnum struct {
	RESERVED ResourceType
	ONDEMAND ResourceType
	HYBRID   ResourceType
}

func GetResourceTypeEnum() ResourceTypeEnum {
	return ResourceTypeEnum{
		RESERVED: ResourceType{
			value: "RESERVED",
		},
		ONDEMAND: ResourceType{
			value: "ONDEMAND",
		},
		HYBRID: ResourceType{
			value: "HYBRID",
		},
	}
}

func (c ResourceType) Value() string {
	return c.value
}

func (c ResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceType) UnmarshalJSON(b []byte) error {
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
