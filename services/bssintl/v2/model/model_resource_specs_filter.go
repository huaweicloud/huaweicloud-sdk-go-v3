package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ResourceSpecsFilter struct {

	// |参数名称：过滤条件键| |参数的约束及描述：必填，仅支持RESOURCE_SPEC|
	Key ResourceSpecsFilterKey `json:"key"`

	// |参数名称：过滤条件值| |参数的约束及描述：必填，过滤条件值|
	Value string `json:"value"`
}

func (o ResourceSpecsFilter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceSpecsFilter struct{}"
	}

	return strings.Join([]string{"ResourceSpecsFilter", string(data)}, " ")
}

type ResourceSpecsFilterKey struct {
	value string
}

type ResourceSpecsFilterKeyEnum struct {
	RESOURCE_SPEC ResourceSpecsFilterKey
}

func GetResourceSpecsFilterKeyEnum() ResourceSpecsFilterKeyEnum {
	return ResourceSpecsFilterKeyEnum{
		RESOURCE_SPEC: ResourceSpecsFilterKey{
			value: "RESOURCE_SPEC",
		},
	}
}

func (c ResourceSpecsFilterKey) Value() string {
	return c.value
}

func (c ResourceSpecsFilterKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceSpecsFilterKey) UnmarshalJSON(b []byte) error {
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
