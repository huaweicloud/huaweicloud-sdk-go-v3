package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateScalingTagInfoRequest Request Object
type CreateScalingTagInfoRequest struct {

	// 资源类型，枚举类：scaling_group_tag。scaling_group_tag表示资源类型为伸缩组。
	ResourceType CreateScalingTagInfoRequestResourceType `json:"resource_type"`

	// 资源ID。
	ResourceId string `json:"resource_id"`

	Body *CreateTagsOption `json:"body,omitempty"`
}

func (o CreateScalingTagInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingTagInfoRequest struct{}"
	}

	return strings.Join([]string{"CreateScalingTagInfoRequest", string(data)}, " ")
}

type CreateScalingTagInfoRequestResourceType struct {
	value string
}

type CreateScalingTagInfoRequestResourceTypeEnum struct {
	SCALING_GROUP_TAG CreateScalingTagInfoRequestResourceType
}

func GetCreateScalingTagInfoRequestResourceTypeEnum() CreateScalingTagInfoRequestResourceTypeEnum {
	return CreateScalingTagInfoRequestResourceTypeEnum{
		SCALING_GROUP_TAG: CreateScalingTagInfoRequestResourceType{
			value: "scaling_group_tag",
		},
	}
}

func (c CreateScalingTagInfoRequestResourceType) Value() string {
	return c.value
}

func (c CreateScalingTagInfoRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateScalingTagInfoRequestResourceType) UnmarshalJSON(b []byte) error {
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
