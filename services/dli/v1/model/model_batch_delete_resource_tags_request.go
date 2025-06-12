package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchDeleteResourceTagsRequest Request Object
type BatchDeleteResourceTagsRequest struct {

	// 资源类型
	ResourceType BatchDeleteResourceTagsRequestResourceType `json:"resource_type"`

	// 资源id
	ResourceId string `json:"resource_id"`

	Body *BatchDeleteResourceTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteResourceTagsRequest", string(data)}, " ")
}

type BatchDeleteResourceTagsRequestResourceType struct {
	value string
}

type BatchDeleteResourceTagsRequestResourceTypeEnum struct {
	DLI_QUEUE                 BatchDeleteResourceTagsRequestResourceType
	DLI_ENHANCED_DATASOURCE   BatchDeleteResourceTagsRequestResourceType
	DLI_DATABASE              BatchDeleteResourceTagsRequestResourceType
	DLI_PACKAGE_RESOURCE      BatchDeleteResourceTagsRequestResourceType
	DLI_FLINK_JOB             BatchDeleteResourceTagsRequestResourceType
	DLI_ELASTIC_RESOURCE_POOL BatchDeleteResourceTagsRequestResourceType
}

func GetBatchDeleteResourceTagsRequestResourceTypeEnum() BatchDeleteResourceTagsRequestResourceTypeEnum {
	return BatchDeleteResourceTagsRequestResourceTypeEnum{
		DLI_QUEUE: BatchDeleteResourceTagsRequestResourceType{
			value: "dli_queue",
		},
		DLI_ENHANCED_DATASOURCE: BatchDeleteResourceTagsRequestResourceType{
			value: "dli_enhanced_datasource",
		},
		DLI_DATABASE: BatchDeleteResourceTagsRequestResourceType{
			value: "dli_database",
		},
		DLI_PACKAGE_RESOURCE: BatchDeleteResourceTagsRequestResourceType{
			value: "dli_package_resource",
		},
		DLI_FLINK_JOB: BatchDeleteResourceTagsRequestResourceType{
			value: "dli_flink_job",
		},
		DLI_ELASTIC_RESOURCE_POOL: BatchDeleteResourceTagsRequestResourceType{
			value: "dli_elastic_resource_pool",
		},
	}
}

func (c BatchDeleteResourceTagsRequestResourceType) Value() string {
	return c.value
}

func (c BatchDeleteResourceTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteResourceTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
