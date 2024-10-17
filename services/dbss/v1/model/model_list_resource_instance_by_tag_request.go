package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResourceInstanceByTagRequest Request Object
type ListResourceInstanceByTagRequest struct {

	// 资源类型。 - auditInstance
	ResourceType ListResourceInstanceByTagRequestResourceType `json:"resource_type"`

	// 查询记录数（action为count时无此参数）如果action为filter默认为1000，limit最多为1000,不能为负数，最小值为1。
	Limit *string `json:"limit,omitempty"`

	// 索引位置，偏移量（action为count时无此参数）从第一条数据偏移offset条数据后开始查询，如果action为filter默认为0（偏移0条数据，表示从第一条数据开始查询）,必须为数字，不能为负数。
	Offset *string `json:"offset,omitempty"`

	Body *ResourceInstanceTagRequest `json:"body,omitempty"`
}

func (o ListResourceInstanceByTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceInstanceByTagRequest struct{}"
	}

	return strings.Join([]string{"ListResourceInstanceByTagRequest", string(data)}, " ")
}

type ListResourceInstanceByTagRequestResourceType struct {
	value string
}

type ListResourceInstanceByTagRequestResourceTypeEnum struct {
	AUDIT_INSTANCE ListResourceInstanceByTagRequestResourceType
}

func GetListResourceInstanceByTagRequestResourceTypeEnum() ListResourceInstanceByTagRequestResourceTypeEnum {
	return ListResourceInstanceByTagRequestResourceTypeEnum{
		AUDIT_INSTANCE: ListResourceInstanceByTagRequestResourceType{
			value: "auditInstance",
		},
	}
}

func (c ListResourceInstanceByTagRequestResourceType) Value() string {
	return c.value
}

func (c ListResourceInstanceByTagRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourceInstanceByTagRequestResourceType) UnmarshalJSON(b []byte) error {
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
