package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// InstancePlatformType 运行平台类型。  应用可以在不同的平台上运行，可选用的平台的类型有以下几种：cce、vmapp。
type InstancePlatformType struct {
	value string
}

type InstancePlatformTypeEnum struct {
	CCE   InstancePlatformType
	VMAPP InstancePlatformType
}

func GetInstancePlatformTypeEnum() InstancePlatformTypeEnum {
	return InstancePlatformTypeEnum{
		CCE: InstancePlatformType{
			value: "cce",
		},
		VMAPP: InstancePlatformType{
			value: "vmapp",
		},
	}
}

func (c InstancePlatformType) Value() string {
	return c.value
}

func (c InstancePlatformType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstancePlatformType) UnmarshalJSON(b []byte) error {
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
