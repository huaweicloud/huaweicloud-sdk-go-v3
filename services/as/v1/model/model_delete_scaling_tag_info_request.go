package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteScalingTagInfoRequest Request Object
type DeleteScalingTagInfoRequest struct {

	// 资源类型，枚举类：scaling_group_tag。scaling_group_tag表示资源类型为伸缩组。
	ResourceType DeleteScalingTagInfoRequestResourceType `json:"resource_type"`

	// 资源ID。
	ResourceId string `json:"resource_id"`

	Body *DeleteTagsOption `json:"body,omitempty"`
}

func (o DeleteScalingTagInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScalingTagInfoRequest struct{}"
	}

	return strings.Join([]string{"DeleteScalingTagInfoRequest", string(data)}, " ")
}

type DeleteScalingTagInfoRequestResourceType struct {
	value string
}

type DeleteScalingTagInfoRequestResourceTypeEnum struct {
	SCALING_GROUP_TAG DeleteScalingTagInfoRequestResourceType
}

func GetDeleteScalingTagInfoRequestResourceTypeEnum() DeleteScalingTagInfoRequestResourceTypeEnum {
	return DeleteScalingTagInfoRequestResourceTypeEnum{
		SCALING_GROUP_TAG: DeleteScalingTagInfoRequestResourceType{
			value: "scaling_group_tag",
		},
	}
}

func (c DeleteScalingTagInfoRequestResourceType) Value() string {
	return c.value
}

func (c DeleteScalingTagInfoRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteScalingTagInfoRequestResourceType) UnmarshalJSON(b []byte) error {
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
