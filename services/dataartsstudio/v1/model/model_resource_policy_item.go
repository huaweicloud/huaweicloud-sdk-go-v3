package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ResourcePolicyItem struct {

	// 资源id
	ResourceId string `json:"resource_id"`

	// 资源名称
	ResourceName string `json:"resource_name"`

	// 资源类型:DATA_CONNECTION,AGENCY
	ResourceType ResourcePolicyItemResourceType `json:"resource_type"`
}

func (o ResourcePolicyItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcePolicyItem struct{}"
	}

	return strings.Join([]string{"ResourcePolicyItem", string(data)}, " ")
}

type ResourcePolicyItemResourceType struct {
	value string
}

type ResourcePolicyItemResourceTypeEnum struct {
	DATA_CONNECTION ResourcePolicyItemResourceType
	AGENCY          ResourcePolicyItemResourceType
}

func GetResourcePolicyItemResourceTypeEnum() ResourcePolicyItemResourceTypeEnum {
	return ResourcePolicyItemResourceTypeEnum{
		DATA_CONNECTION: ResourcePolicyItemResourceType{
			value: "DATA_CONNECTION",
		},
		AGENCY: ResourcePolicyItemResourceType{
			value: "AGENCY",
		},
	}
}

func (c ResourcePolicyItemResourceType) Value() string {
	return c.value
}

func (c ResourcePolicyItemResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourcePolicyItemResourceType) UnmarshalJSON(b []byte) error {
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
