package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTagsForResourceTypeRequest Request Object
type ListTagsForResourceTypeRequest struct {

	// 资源类型
	ResourceType ListTagsForResourceTypeRequestResourceType `json:"resource_type"`

	// 查询记录数默认为1000，limit最多为1000,不能为负数，最小值为1
	Limit *int32 `json:"limit,omitempty"`

	// 索引位置，从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询）,必须为数字，不能为负数
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListTagsForResourceTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsForResourceTypeRequest struct{}"
	}

	return strings.Join([]string{"ListTagsForResourceTypeRequest", string(data)}, " ")
}

type ListTagsForResourceTypeRequestResourceType struct {
	value string
}

type ListTagsForResourceTypeRequestResourceTypeEnum struct {
	CONFIGPOLICY_ASSIGNMENTS ListTagsForResourceTypeRequestResourceType
}

func GetListTagsForResourceTypeRequestResourceTypeEnum() ListTagsForResourceTypeRequestResourceTypeEnum {
	return ListTagsForResourceTypeRequestResourceTypeEnum{
		CONFIGPOLICY_ASSIGNMENTS: ListTagsForResourceTypeRequestResourceType{
			value: "config:policyAssignments",
		},
	}
}

func (c ListTagsForResourceTypeRequestResourceType) Value() string {
	return c.value
}

func (c ListTagsForResourceTypeRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTagsForResourceTypeRequestResourceType) UnmarshalJSON(b []byte) error {
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
