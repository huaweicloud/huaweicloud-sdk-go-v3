package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CountResourcesByTagRequest Request Object
type CountResourcesByTagRequest struct {

	// 资源类型
	ResourceType CountResourcesByTagRequestResourceType `json:"resource_type"`

	Body *ResourceInstancesReq `json:"body,omitempty"`
}

func (o CountResourcesByTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountResourcesByTagRequest struct{}"
	}

	return strings.Join([]string{"CountResourcesByTagRequest", string(data)}, " ")
}

type CountResourcesByTagRequestResourceType struct {
	value string
}

type CountResourcesByTagRequestResourceTypeEnum struct {
	CONFIGPOLICY_ASSIGNMENTS CountResourcesByTagRequestResourceType
}

func GetCountResourcesByTagRequestResourceTypeEnum() CountResourcesByTagRequestResourceTypeEnum {
	return CountResourcesByTagRequestResourceTypeEnum{
		CONFIGPOLICY_ASSIGNMENTS: CountResourcesByTagRequestResourceType{
			value: "config:policyAssignments",
		},
	}
}

func (c CountResourcesByTagRequestResourceType) Value() string {
	return c.value
}

func (c CountResourcesByTagRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CountResourcesByTagRequestResourceType) UnmarshalJSON(b []byte) error {
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
