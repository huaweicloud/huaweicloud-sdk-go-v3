package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTagsRequest Request Object
type ListTagsRequest struct {

	// 资源类型: - cloud-connection: 云连接 - bandwidth-package: 带宽包
	ResourceType ListTagsRequestResourceType `json:"resource_type"`

	// 资源ID
	ResourceId string `json:"resource_id"`
}

func (o ListTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsRequest struct{}"
	}

	return strings.Join([]string{"ListTagsRequest", string(data)}, " ")
}

type ListTagsRequestResourceType struct {
	value string
}

type ListTagsRequestResourceTypeEnum struct {
	CLOUD_CONNECTION  ListTagsRequestResourceType
	BANDWIDTH_PACKAGE ListTagsRequestResourceType
}

func GetListTagsRequestResourceTypeEnum() ListTagsRequestResourceTypeEnum {
	return ListTagsRequestResourceTypeEnum{
		CLOUD_CONNECTION: ListTagsRequestResourceType{
			value: "cloud-connection",
		},
		BANDWIDTH_PACKAGE: ListTagsRequestResourceType{
			value: "bandwidth-package",
		},
	}
}

func (c ListTagsRequestResourceType) Value() string {
	return c.value
}

func (c ListTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
