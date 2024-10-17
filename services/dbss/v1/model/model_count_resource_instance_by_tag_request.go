package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CountResourceInstanceByTagRequest Request Object
type CountResourceInstanceByTagRequest struct {

	// 资源类型。 - auditInstance
	ResourceType CountResourceInstanceByTagRequestResourceType `json:"resource_type"`

	Body *ResourceInstanceTagRequest `json:"body,omitempty"`
}

func (o CountResourceInstanceByTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountResourceInstanceByTagRequest struct{}"
	}

	return strings.Join([]string{"CountResourceInstanceByTagRequest", string(data)}, " ")
}

type CountResourceInstanceByTagRequestResourceType struct {
	value string
}

type CountResourceInstanceByTagRequestResourceTypeEnum struct {
	AUDIT_INSTANCE CountResourceInstanceByTagRequestResourceType
}

func GetCountResourceInstanceByTagRequestResourceTypeEnum() CountResourceInstanceByTagRequestResourceTypeEnum {
	return CountResourceInstanceByTagRequestResourceTypeEnum{
		AUDIT_INSTANCE: CountResourceInstanceByTagRequestResourceType{
			value: "auditInstance",
		},
	}
}

func (c CountResourceInstanceByTagRequestResourceType) Value() string {
	return c.value
}

func (c CountResourceInstanceByTagRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CountResourceInstanceByTagRequestResourceType) UnmarshalJSON(b []byte) error {
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
