package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListResourceByFilterTagRequest struct {

	// 资源类型: - cc: 云连接 - bwp: 带宽包
	ResourceType ListResourceByFilterTagRequestResourceType `json:"resource_type"`

	Body *ListResourceByFilterTagRequestBody `json:"body,omitempty"`
}

func (o ListResourceByFilterTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceByFilterTagRequest struct{}"
	}

	return strings.Join([]string{"ListResourceByFilterTagRequest", string(data)}, " ")
}

type ListResourceByFilterTagRequestResourceType struct {
	value string
}

type ListResourceByFilterTagRequestResourceTypeEnum struct {
	CC  ListResourceByFilterTagRequestResourceType
	BWP ListResourceByFilterTagRequestResourceType
}

func GetListResourceByFilterTagRequestResourceTypeEnum() ListResourceByFilterTagRequestResourceTypeEnum {
	return ListResourceByFilterTagRequestResourceTypeEnum{
		CC: ListResourceByFilterTagRequestResourceType{
			value: "cc",
		},
		BWP: ListResourceByFilterTagRequestResourceType{
			value: "bwp",
		},
	}
}

func (c ListResourceByFilterTagRequestResourceType) Value() string {
	return c.value
}

func (c ListResourceByFilterTagRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourceByFilterTagRequestResourceType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
