package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowResourceTagsRequest Request Object
type ShowResourceTagsRequest struct {

	// 资源类型
	ResourceType ShowResourceTagsRequestResourceType `json:"resource_type"`

	// 资源id
	ResourceId string `json:"resource_id"`

	// 查询记录数
	Limit *int32 `json:"limit,omitempty"`

	// 索引位置偏移量
	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceTagsRequest", string(data)}, " ")
}

type ShowResourceTagsRequestResourceType struct {
	value string
}

type ShowResourceTagsRequestResourceTypeEnum struct {
	DLI_QUEUE                 ShowResourceTagsRequestResourceType
	DLI_ENHANCED_DATASOURCE   ShowResourceTagsRequestResourceType
	DLI_DATABASE              ShowResourceTagsRequestResourceType
	DLI_PACKAGE_RESOURCE      ShowResourceTagsRequestResourceType
	DLI_FLINK_JOB             ShowResourceTagsRequestResourceType
	DLI_ELASTIC_RESOURCE_POOL ShowResourceTagsRequestResourceType
}

func GetShowResourceTagsRequestResourceTypeEnum() ShowResourceTagsRequestResourceTypeEnum {
	return ShowResourceTagsRequestResourceTypeEnum{
		DLI_QUEUE: ShowResourceTagsRequestResourceType{
			value: "dli_queue",
		},
		DLI_ENHANCED_DATASOURCE: ShowResourceTagsRequestResourceType{
			value: "dli_enhanced_datasource",
		},
		DLI_DATABASE: ShowResourceTagsRequestResourceType{
			value: "dli_database",
		},
		DLI_PACKAGE_RESOURCE: ShowResourceTagsRequestResourceType{
			value: "dli_package_resource",
		},
		DLI_FLINK_JOB: ShowResourceTagsRequestResourceType{
			value: "dli_flink_job",
		},
		DLI_ELASTIC_RESOURCE_POOL: ShowResourceTagsRequestResourceType{
			value: "dli_elastic_resource_pool",
		},
	}
}

func (c ShowResourceTagsRequestResourceType) Value() string {
	return c.value
}

func (c ShowResourceTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowResourceTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
