package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchAddOrRemoveResourceInstanceRequest Request Object
type BatchAddOrRemoveResourceInstanceRequest struct {

	// 资源类型。  - endpoint_service：终端节点服务  - endpoint：终端节点
	ResourceType BatchAddOrRemoveResourceInstanceRequestResourceType `json:"resource_type"`

	// 资源ID，Endpoint Service ID或Endpoint ID。
	ResourceId string `json:"resource_id"`

	// 发送的实体的MIME类型。推荐用户默认使用application/json， 如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。
	ContentType *string `json:"Content-Type,omitempty"`

	Body *BatchAddOrRemoveResourceInstanceRequestBody `json:"body,omitempty"`
}

func (o BatchAddOrRemoveResourceInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddOrRemoveResourceInstanceRequest struct{}"
	}

	return strings.Join([]string{"BatchAddOrRemoveResourceInstanceRequest", string(data)}, " ")
}

type BatchAddOrRemoveResourceInstanceRequestResourceType struct {
	value string
}

type BatchAddOrRemoveResourceInstanceRequestResourceTypeEnum struct {
	ENDPOINT_SERVICE BatchAddOrRemoveResourceInstanceRequestResourceType
	ENDPOINT         BatchAddOrRemoveResourceInstanceRequestResourceType
}

func GetBatchAddOrRemoveResourceInstanceRequestResourceTypeEnum() BatchAddOrRemoveResourceInstanceRequestResourceTypeEnum {
	return BatchAddOrRemoveResourceInstanceRequestResourceTypeEnum{
		ENDPOINT_SERVICE: BatchAddOrRemoveResourceInstanceRequestResourceType{
			value: "endpoint_service",
		},
		ENDPOINT: BatchAddOrRemoveResourceInstanceRequestResourceType{
			value: "endpoint",
		},
	}
}

func (c BatchAddOrRemoveResourceInstanceRequestResourceType) Value() string {
	return c.value
}

func (c BatchAddOrRemoveResourceInstanceRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchAddOrRemoveResourceInstanceRequestResourceType) UnmarshalJSON(b []byte) error {
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
