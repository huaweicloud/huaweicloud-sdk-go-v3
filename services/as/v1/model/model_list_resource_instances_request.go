package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResourceInstancesRequest Request Object
type ListResourceInstancesRequest struct {

	// 资源类型，枚举类：SCALING_GROUP_TAG。
	ResourceType ListResourceInstancesRequestResourceType `json:"resource_type"`

	Body *QueryTagsOption `json:"body,omitempty"`
}

func (o ListResourceInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListResourceInstancesRequest", string(data)}, " ")
}

type ListResourceInstancesRequestResourceType struct {
	value string
}

type ListResourceInstancesRequestResourceTypeEnum struct {
	SCALING_GROUP_TAG ListResourceInstancesRequestResourceType
}

func GetListResourceInstancesRequestResourceTypeEnum() ListResourceInstancesRequestResourceTypeEnum {
	return ListResourceInstancesRequestResourceTypeEnum{
		SCALING_GROUP_TAG: ListResourceInstancesRequestResourceType{
			value: "scaling_group_tag",
		},
	}
}

func (c ListResourceInstancesRequestResourceType) Value() string {
	return c.value
}

func (c ListResourceInstancesRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourceInstancesRequestResourceType) UnmarshalJSON(b []byte) error {
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
