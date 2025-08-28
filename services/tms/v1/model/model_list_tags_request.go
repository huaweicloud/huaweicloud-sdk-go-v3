package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTagsRequest Request Object
type ListTagsRequest struct {

	// 资源类型
	ResourceTypes string `json:"resource_types"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 标签类型
	TagTypes ListTagsRequestTagTypes `json:"tag_types"`
}

func (o ListTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsRequest struct{}"
	}

	return strings.Join([]string{"ListTagsRequest", string(data)}, " ")
}

type ListTagsRequestTagTypes struct {
	value string
}

type ListTagsRequestTagTypesEnum struct {
	RESOURCE ListTagsRequestTagTypes
}

func GetListTagsRequestTagTypesEnum() ListTagsRequestTagTypesEnum {
	return ListTagsRequestTagTypesEnum{
		RESOURCE: ListTagsRequestTagTypes{
			value: "resource",
		},
	}
}

func (c ListTagsRequestTagTypes) Value() string {
	return c.value
}

func (c ListTagsRequestTagTypes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTagsRequestTagTypes) UnmarshalJSON(b []byte) error {
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
