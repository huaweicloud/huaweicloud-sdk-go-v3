package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResourcesByTagRequest Request Object
type ListResourcesByTagRequest struct {

	// 资源类型
	ResourceType ListResourcesByTagRequestResourceType `json:"resource_type"`

	// 查询记录数（action为count时无此参数）如果action为filter默认为1000，limit最多为1000,不能为负数，最小值为1
	Limit *int32 `json:"limit,omitempty"`

	// 索引位置，偏移量（action为count时无此参数）从第一条数据偏移offset条数据后开始查询，如果action为filter默认为0（偏移0条数据，表示从第一条数据开始查询）,必须为数字，不能为负数
	Offset *int32 `json:"offset,omitempty"`

	Body *ResourceInstancesReq `json:"body,omitempty"`
}

func (o ListResourcesByTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesByTagRequest struct{}"
	}

	return strings.Join([]string{"ListResourcesByTagRequest", string(data)}, " ")
}

type ListResourcesByTagRequestResourceType struct {
	value string
}

type ListResourcesByTagRequestResourceTypeEnum struct {
	CONFIGPOLICY_ASSIGNMENTS ListResourcesByTagRequestResourceType
}

func GetListResourcesByTagRequestResourceTypeEnum() ListResourcesByTagRequestResourceTypeEnum {
	return ListResourcesByTagRequestResourceTypeEnum{
		CONFIGPOLICY_ASSIGNMENTS: ListResourcesByTagRequestResourceType{
			value: "config:policyAssignments",
		},
	}
}

func (c ListResourcesByTagRequestResourceType) Value() string {
	return c.value
}

func (c ListResourcesByTagRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourcesByTagRequestResourceType) UnmarshalJSON(b []byte) error {
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
