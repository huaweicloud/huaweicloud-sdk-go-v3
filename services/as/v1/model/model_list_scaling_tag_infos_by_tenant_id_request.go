package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListScalingTagInfosByTenantIdRequest Request Object
type ListScalingTagInfosByTenantIdRequest struct {

	// 资源类型，枚举类：scaling_group_tag。scaling_group_tag表示资源类型为伸缩组。
	ResourceType ListScalingTagInfosByTenantIdRequestResourceType `json:"resource_type"`
}

func (o ListScalingTagInfosByTenantIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingTagInfosByTenantIdRequest struct{}"
	}

	return strings.Join([]string{"ListScalingTagInfosByTenantIdRequest", string(data)}, " ")
}

type ListScalingTagInfosByTenantIdRequestResourceType struct {
	value string
}

type ListScalingTagInfosByTenantIdRequestResourceTypeEnum struct {
	SCALING_GROUP_TAG ListScalingTagInfosByTenantIdRequestResourceType
}

func GetListScalingTagInfosByTenantIdRequestResourceTypeEnum() ListScalingTagInfosByTenantIdRequestResourceTypeEnum {
	return ListScalingTagInfosByTenantIdRequestResourceTypeEnum{
		SCALING_GROUP_TAG: ListScalingTagInfosByTenantIdRequestResourceType{
			value: "scaling_group_tag",
		},
	}
}

func (c ListScalingTagInfosByTenantIdRequestResourceType) Value() string {
	return c.value
}

func (c ListScalingTagInfosByTenantIdRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListScalingTagInfosByTenantIdRequestResourceType) UnmarshalJSON(b []byte) error {
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
