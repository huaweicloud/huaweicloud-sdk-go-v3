package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchCreateResourceTagsRequest Request Object
type BatchCreateResourceTagsRequest struct {

	// 资源类型
	ResourceType BatchCreateResourceTagsRequestResourceType `json:"resource_type"`

	// 资源id
	ResourceId string `json:"resource_id"`

	Body *BatchCreateResourceTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateResourceTagsRequest", string(data)}, " ")
}

type BatchCreateResourceTagsRequestResourceType struct {
	value string
}

type BatchCreateResourceTagsRequestResourceTypeEnum struct {
	DLI_QUEUE                 BatchCreateResourceTagsRequestResourceType
	DLI_ENHANCED_DATASOURCE   BatchCreateResourceTagsRequestResourceType
	DLI_DATABASE              BatchCreateResourceTagsRequestResourceType
	DLI_PACKAGE_RESOURCE      BatchCreateResourceTagsRequestResourceType
	DLI_FLINK_JOB             BatchCreateResourceTagsRequestResourceType
	DLI_ELASTIC_RESOURCE_POOL BatchCreateResourceTagsRequestResourceType
}

func GetBatchCreateResourceTagsRequestResourceTypeEnum() BatchCreateResourceTagsRequestResourceTypeEnum {
	return BatchCreateResourceTagsRequestResourceTypeEnum{
		DLI_QUEUE: BatchCreateResourceTagsRequestResourceType{
			value: "dli_queue",
		},
		DLI_ENHANCED_DATASOURCE: BatchCreateResourceTagsRequestResourceType{
			value: "dli_enhanced_datasource",
		},
		DLI_DATABASE: BatchCreateResourceTagsRequestResourceType{
			value: "dli_database",
		},
		DLI_PACKAGE_RESOURCE: BatchCreateResourceTagsRequestResourceType{
			value: "dli_package_resource",
		},
		DLI_FLINK_JOB: BatchCreateResourceTagsRequestResourceType{
			value: "dli_flink_job",
		},
		DLI_ELASTIC_RESOURCE_POOL: BatchCreateResourceTagsRequestResourceType{
			value: "dli_elastic_resource_pool",
		},
	}
}

func (c BatchCreateResourceTagsRequestResourceType) Value() string {
	return c.value
}

func (c BatchCreateResourceTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateResourceTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
