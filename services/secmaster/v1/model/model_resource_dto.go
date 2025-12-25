package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ResourceDto struct {

	// 是否使能： active, inactive
	Enable ResourceDtoEnable `json:"enable"`

	// 区域ID
	RegionId string `json:"region_id"`
}

func (o ResourceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceDto struct{}"
	}

	return strings.Join([]string{"ResourceDto", string(data)}, " ")
}

type ResourceDtoEnable struct {
	value string
}

type ResourceDtoEnableEnum struct {
	INACTIVE ResourceDtoEnable
	ACTIVE   ResourceDtoEnable
}

func GetResourceDtoEnableEnum() ResourceDtoEnableEnum {
	return ResourceDtoEnableEnum{
		INACTIVE: ResourceDtoEnable{
			value: "inactive",
		},
		ACTIVE: ResourceDtoEnable{
			value: "active",
		},
	}
}

func (c ResourceDtoEnable) Value() string {
	return c.value
}

func (c ResourceDtoEnable) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceDtoEnable) UnmarshalJSON(b []byte) error {
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
