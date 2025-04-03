package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResourceLevel 产品层级跨纬规则创建时需要指明为产品层级规则，resource_level取值为product即为产品层级跨纬规则，不填或者取值为dimension则为旧的规则类型
type ResourceLevel struct {
	value string
}

type ResourceLevelEnum struct {
	PRODUCT   ResourceLevel
	DIMENSION ResourceLevel
}

func GetResourceLevelEnum() ResourceLevelEnum {
	return ResourceLevelEnum{
		PRODUCT: ResourceLevel{
			value: "product",
		},
		DIMENSION: ResourceLevel{
			value: "dimension",
		},
	}
}

func (c ResourceLevel) Value() string {
	return c.value
}

func (c ResourceLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceLevel) UnmarshalJSON(b []byte) error {
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
