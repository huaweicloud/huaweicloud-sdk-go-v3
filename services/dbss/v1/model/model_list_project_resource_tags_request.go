package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListProjectResourceTagsRequest Request Object
type ListProjectResourceTagsRequest struct {

	// 资源类型。 - auditInstance
	ResourceType ListProjectResourceTagsRequestResourceType `json:"resource_type"`
}

func (o ListProjectResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListProjectResourceTagsRequest", string(data)}, " ")
}

type ListProjectResourceTagsRequestResourceType struct {
	value string
}

type ListProjectResourceTagsRequestResourceTypeEnum struct {
	AUDIT_INSTANCE ListProjectResourceTagsRequestResourceType
}

func GetListProjectResourceTagsRequestResourceTypeEnum() ListProjectResourceTagsRequestResourceTypeEnum {
	return ListProjectResourceTagsRequestResourceTypeEnum{
		AUDIT_INSTANCE: ListProjectResourceTagsRequestResourceType{
			value: "auditInstance",
		},
	}
}

func (c ListProjectResourceTagsRequestResourceType) Value() string {
	return c.value
}

func (c ListProjectResourceTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProjectResourceTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
