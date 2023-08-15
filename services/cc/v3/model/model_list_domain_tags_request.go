package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDomainTagsRequest Request Object
type ListDomainTagsRequest struct {

	// 资源类型: - cloud-connection: 云连接 - bandwidth-package: 带宽包
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
	CLOUD_CONNECTION  ListDomainTagsRequestResourceType
	BANDWIDTH_PACKAGE ListDomainTagsRequestResourceType
}

func GetListDomainTagsRequestResourceTypeEnum() ListDomainTagsRequestResourceTypeEnum {
	return ListDomainTagsRequestResourceTypeEnum{
		CLOUD_CONNECTION: ListDomainTagsRequestResourceType{
			value: "cloud-connection",
		},
		BANDWIDTH_PACKAGE: ListDomainTagsRequestResourceType{
			value: "bandwidth-package",
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
