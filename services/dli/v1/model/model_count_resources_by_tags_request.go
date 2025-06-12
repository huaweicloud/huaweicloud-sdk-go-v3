package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CountResourcesByTagsRequest Request Object
type CountResourcesByTagsRequest struct {

	// 资源类型
	ResourceType CountResourcesByTagsRequestResourceType `json:"resource_type"`

	Body *CountResourcesByTagsRequestBody `json:"body,omitempty"`
}

func (o CountResourcesByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountResourcesByTagsRequest struct{}"
	}

	return strings.Join([]string{"CountResourcesByTagsRequest", string(data)}, " ")
}

type CountResourcesByTagsRequestResourceType struct {
	value string
}

type CountResourcesByTagsRequestResourceTypeEnum struct {
	DLI_QUEUE                 CountResourcesByTagsRequestResourceType
	DLI_ENHANCED_DATASOURCE   CountResourcesByTagsRequestResourceType
	DLI_DATABASE              CountResourcesByTagsRequestResourceType
	DLI_PACKAGE_RESOURCE      CountResourcesByTagsRequestResourceType
	DLI_FLINK_JOB             CountResourcesByTagsRequestResourceType
	DLI_ELASTIC_RESOURCE_POOL CountResourcesByTagsRequestResourceType
}

func GetCountResourcesByTagsRequestResourceTypeEnum() CountResourcesByTagsRequestResourceTypeEnum {
	return CountResourcesByTagsRequestResourceTypeEnum{
		DLI_QUEUE: CountResourcesByTagsRequestResourceType{
			value: "dli_queue",
		},
		DLI_ENHANCED_DATASOURCE: CountResourcesByTagsRequestResourceType{
			value: "dli_enhanced_datasource",
		},
		DLI_DATABASE: CountResourcesByTagsRequestResourceType{
			value: "dli_database",
		},
		DLI_PACKAGE_RESOURCE: CountResourcesByTagsRequestResourceType{
			value: "dli_package_resource",
		},
		DLI_FLINK_JOB: CountResourcesByTagsRequestResourceType{
			value: "dli_flink_job",
		},
		DLI_ELASTIC_RESOURCE_POOL: CountResourcesByTagsRequestResourceType{
			value: "dli_elastic_resource_pool",
		},
	}
}

func (c CountResourcesByTagsRequestResourceType) Value() string {
	return c.value
}

func (c CountResourcesByTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CountResourcesByTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
