package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResourcesByTagsRequest Request Object
type ListResourcesByTagsRequest struct {

	// 资源类型
	ResourceType ListResourcesByTagsRequestResourceType `json:"resource_type"`

	// 查询记录数
	Limit *int32 `json:"limit,omitempty"`

	// 索引位置偏移量
	Offset *int32 `json:"offset,omitempty"`

	Body *ListResourcesByTagsRequestBody `json:"body,omitempty"`
}

func (o ListResourcesByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListResourcesByTagsRequest", string(data)}, " ")
}

type ListResourcesByTagsRequestResourceType struct {
	value string
}

type ListResourcesByTagsRequestResourceTypeEnum struct {
	DLI_QUEUE                 ListResourcesByTagsRequestResourceType
	DLI_ENHANCED_DATASOURCE   ListResourcesByTagsRequestResourceType
	DLI_DATABASE              ListResourcesByTagsRequestResourceType
	DLI_PACKAGE_RESOURCE      ListResourcesByTagsRequestResourceType
	DLI_FLINK_JOB             ListResourcesByTagsRequestResourceType
	DLI_ELASTIC_RESOURCE_POOL ListResourcesByTagsRequestResourceType
}

func GetListResourcesByTagsRequestResourceTypeEnum() ListResourcesByTagsRequestResourceTypeEnum {
	return ListResourcesByTagsRequestResourceTypeEnum{
		DLI_QUEUE: ListResourcesByTagsRequestResourceType{
			value: "dli_queue",
		},
		DLI_ENHANCED_DATASOURCE: ListResourcesByTagsRequestResourceType{
			value: "dli_enhanced_datasource",
		},
		DLI_DATABASE: ListResourcesByTagsRequestResourceType{
			value: "dli_database",
		},
		DLI_PACKAGE_RESOURCE: ListResourcesByTagsRequestResourceType{
			value: "dli_package_resource",
		},
		DLI_FLINK_JOB: ListResourcesByTagsRequestResourceType{
			value: "dli_flink_job",
		},
		DLI_ELASTIC_RESOURCE_POOL: ListResourcesByTagsRequestResourceType{
			value: "dli_elastic_resource_pool",
		},
	}
}

func (c ListResourcesByTagsRequestResourceType) Value() string {
	return c.value
}

func (c ListResourcesByTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourcesByTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
