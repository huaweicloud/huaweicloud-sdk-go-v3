package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchCreateDeleteTagsRequest Request Object
type BatchCreateDeleteTagsRequest struct {

	// 资源ID
	ResourceId string `json:"resource_id"`

	// 资源类型: - cloud-connection: 云连接 - bandwidth-package: 带宽包
	ResourceType BatchCreateDeleteTagsRequestResourceType `json:"resource_type"`

	Body *Tags `json:"body,omitempty"`
}

func (o BatchCreateDeleteTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDeleteTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateDeleteTagsRequest", string(data)}, " ")
}

type BatchCreateDeleteTagsRequestResourceType struct {
	value string
}

type BatchCreateDeleteTagsRequestResourceTypeEnum struct {
	CLOUD_CONNECTION  BatchCreateDeleteTagsRequestResourceType
	BANDWIDTH_PACKAGE BatchCreateDeleteTagsRequestResourceType
}

func GetBatchCreateDeleteTagsRequestResourceTypeEnum() BatchCreateDeleteTagsRequestResourceTypeEnum {
	return BatchCreateDeleteTagsRequestResourceTypeEnum{
		CLOUD_CONNECTION: BatchCreateDeleteTagsRequestResourceType{
			value: "cloud-connection",
		},
		BANDWIDTH_PACKAGE: BatchCreateDeleteTagsRequestResourceType{
			value: "bandwidth-package",
		},
	}
}

func (c BatchCreateDeleteTagsRequestResourceType) Value() string {
	return c.value
}

func (c BatchCreateDeleteTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateDeleteTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
