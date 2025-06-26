package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstanceResourceInstancesRequest Request Object
type ListInstanceResourceInstancesRequest struct {

	// 资源类型，支持的资源类型为：instances
	ResourceType ListInstanceResourceInstancesRequestResourceType `json:"resource_type"`

	// 起始索引，默认值为0。**注意：offset和limit参数需要配套使用。**
	Offset *int32 `json:"offset,omitempty"`

	// 返回条数，默认为200，最小值为1，最大值为200。**注意：offset和limit参数需要配套使用。**
	Limit *int32 `json:"limit,omitempty"`

	Body *ListResourceInstancesRequestBody `json:"body,omitempty"`
}

func (o ListInstanceResourceInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceResourceInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceResourceInstancesRequest", string(data)}, " ")
}

type ListInstanceResourceInstancesRequestResourceType struct {
	value string
}

type ListInstanceResourceInstancesRequestResourceTypeEnum struct {
	INSTANCES ListInstanceResourceInstancesRequestResourceType
}

func GetListInstanceResourceInstancesRequestResourceTypeEnum() ListInstanceResourceInstancesRequestResourceTypeEnum {
	return ListInstanceResourceInstancesRequestResourceTypeEnum{
		INSTANCES: ListInstanceResourceInstancesRequestResourceType{
			value: "instances",
		},
	}
}

func (c ListInstanceResourceInstancesRequestResourceType) Value() string {
	return c.value
}

func (c ListInstanceResourceInstancesRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceResourceInstancesRequestResourceType) UnmarshalJSON(b []byte) error {
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
