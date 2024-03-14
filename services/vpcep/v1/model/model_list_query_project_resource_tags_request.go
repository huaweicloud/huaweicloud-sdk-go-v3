package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListQueryProjectResourceTagsRequest Request Object
type ListQueryProjectResourceTagsRequest struct {

	// 资源类型。  - endpoint_service：终端节点服务  - endpoint：终端节点
	ResourceType ListQueryProjectResourceTagsRequestResourceType `json:"resource_type"`

	// 发送的实体的MIME类型。推荐用户默认使用application/json， 如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。
	ContentType *string `json:"Content-Type,omitempty"`
}

func (o ListQueryProjectResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueryProjectResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListQueryProjectResourceTagsRequest", string(data)}, " ")
}

type ListQueryProjectResourceTagsRequestResourceType struct {
	value string
}

type ListQueryProjectResourceTagsRequestResourceTypeEnum struct {
	ENDPOINT_SERVICE ListQueryProjectResourceTagsRequestResourceType
	ENDPOINT         ListQueryProjectResourceTagsRequestResourceType
}

func GetListQueryProjectResourceTagsRequestResourceTypeEnum() ListQueryProjectResourceTagsRequestResourceTypeEnum {
	return ListQueryProjectResourceTagsRequestResourceTypeEnum{
		ENDPOINT_SERVICE: ListQueryProjectResourceTagsRequestResourceType{
			value: "endpoint_service",
		},
		ENDPOINT: ListQueryProjectResourceTagsRequestResourceType{
			value: "endpoint",
		},
	}
}

func (c ListQueryProjectResourceTagsRequestResourceType) Value() string {
	return c.value
}

func (c ListQueryProjectResourceTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListQueryProjectResourceTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
