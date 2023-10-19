package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResourceType 带宽包实例绑定的资源类型。
type ResourceType struct {

	// 带宽包实例绑定的资源类型。 cloud_connection: 云连接实例。
	ResourceType ResourceTypeResourceType `json:"resource_type"`
}

func (o ResourceType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceType struct{}"
	}

	return strings.Join([]string{"ResourceType", string(data)}, " ")
}

type ResourceTypeResourceType struct {
	value string
}

type ResourceTypeResourceTypeEnum struct {
	CLOUD_CONNECTION ResourceTypeResourceType
}

func GetResourceTypeResourceTypeEnum() ResourceTypeResourceTypeEnum {
	return ResourceTypeResourceTypeEnum{
		CLOUD_CONNECTION: ResourceTypeResourceType{
			value: "cloud_connection",
		},
	}
}

func (c ResourceTypeResourceType) Value() string {
	return c.value
}

func (c ResourceTypeResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceTypeResourceType) UnmarshalJSON(b []byte) error {
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
