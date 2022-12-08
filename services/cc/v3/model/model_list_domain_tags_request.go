package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListDomainTagsRequest struct {

	// 资源类型: - cc: 云连接 - bwp: 带宽包
	ResourceType ListDomainTagsRequestResourceType `json:"resource_type"`
}

func (o ListDomainTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainTagsRequest struct{}"
	}

	return strings.Join([]string{"ListDomainTagsRequest", string(data)}, " ")
}

type ListDomainTagsRequestResourceType struct {
	value string
}

type ListDomainTagsRequestResourceTypeEnum struct {
	CC  ListDomainTagsRequestResourceType
	BWP ListDomainTagsRequestResourceType
}

func GetListDomainTagsRequestResourceTypeEnum() ListDomainTagsRequestResourceTypeEnum {
	return ListDomainTagsRequestResourceTypeEnum{
		CC: ListDomainTagsRequestResourceType{
			value: "cc",
		},
		BWP: ListDomainTagsRequestResourceType{
			value: "bwp",
		},
	}
}

func (c ListDomainTagsRequestResourceType) Value() string {
	return c.value
}

func (c ListDomainTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDomainTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
