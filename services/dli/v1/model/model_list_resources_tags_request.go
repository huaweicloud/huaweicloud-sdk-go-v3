package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResourcesTagsRequest Request Object
type ListResourcesTagsRequest struct {

	// 资源类型
	ResourceType ListResourcesTagsRequestResourceType `json:"resource_type"`
}

func (o ListResourcesTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesTagsRequest struct{}"
	}

	return strings.Join([]string{"ListResourcesTagsRequest", string(data)}, " ")
}

type ListResourcesTagsRequestResourceType struct {
	value string
}

type ListResourcesTagsRequestResourceTypeEnum struct {
	DLI_QUEUE                 ListResourcesTagsRequestResourceType
	DLI_ENHANCED_DATASOURCE   ListResourcesTagsRequestResourceType
	DLI_DATABASE              ListResourcesTagsRequestResourceType
	DLI_PACKAGE_RESOURCE      ListResourcesTagsRequestResourceType
	DLI_FLINK_JOB             ListResourcesTagsRequestResourceType
	DLI_ELASTIC_RESOURCE_POOL ListResourcesTagsRequestResourceType
}

func GetListResourcesTagsRequestResourceTypeEnum() ListResourcesTagsRequestResourceTypeEnum {
	return ListResourcesTagsRequestResourceTypeEnum{
		DLI_QUEUE: ListResourcesTagsRequestResourceType{
			value: "dli_queue",
		},
		DLI_ENHANCED_DATASOURCE: ListResourcesTagsRequestResourceType{
			value: "dli_enhanced_datasource",
		},
		DLI_DATABASE: ListResourcesTagsRequestResourceType{
			value: "dli_database",
		},
		DLI_PACKAGE_RESOURCE: ListResourcesTagsRequestResourceType{
			value: "dli_package_resource",
		},
		DLI_FLINK_JOB: ListResourcesTagsRequestResourceType{
			value: "dli_flink_job",
		},
		DLI_ELASTIC_RESOURCE_POOL: ListResourcesTagsRequestResourceType{
			value: "dli_elastic_resource_pool",
		},
	}
}

func (c ListResourcesTagsRequestResourceType) Value() string {
	return c.value
}

func (c ListResourcesTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourcesTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
