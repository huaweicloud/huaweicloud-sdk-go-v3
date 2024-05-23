package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTagsForResourceRequest Request Object
type ListTagsForResourceRequest struct {

	// 资源类型
	ResourceType ListTagsForResourceRequestResourceType `json:"resource_type"`

	// 资源ID
	ResourceId string `json:"resource_id"`
}

func (o ListTagsForResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsForResourceRequest struct{}"
	}

	return strings.Join([]string{"ListTagsForResourceRequest", string(data)}, " ")
}

type ListTagsForResourceRequestResourceType struct {
	value string
}

type ListTagsForResourceRequestResourceTypeEnum struct {
	CONFIGPOLICY_ASSIGNMENTS ListTagsForResourceRequestResourceType
}

func GetListTagsForResourceRequestResourceTypeEnum() ListTagsForResourceRequestResourceTypeEnum {
	return ListTagsForResourceRequestResourceTypeEnum{
		CONFIGPOLICY_ASSIGNMENTS: ListTagsForResourceRequestResourceType{
			value: "config:policyAssignments",
		},
	}
}

func (c ListTagsForResourceRequestResourceType) Value() string {
	return c.value
}

func (c ListTagsForResourceRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTagsForResourceRequestResourceType) UnmarshalJSON(b []byte) error {
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
