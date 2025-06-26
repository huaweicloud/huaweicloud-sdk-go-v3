package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowInstanceResourceInstancesCountRequest Request Object
type ShowInstanceResourceInstancesCountRequest struct {

	// 资源类型，支持的资源类型为：instances
	ResourceType ShowInstanceResourceInstancesCountRequestResourceType `json:"resource_type"`

	Body *ListResourceInstancesRequestBody `json:"body,omitempty"`
}

func (o ShowInstanceResourceInstancesCountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceResourceInstancesCountRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceResourceInstancesCountRequest", string(data)}, " ")
}

type ShowInstanceResourceInstancesCountRequestResourceType struct {
	value string
}

type ShowInstanceResourceInstancesCountRequestResourceTypeEnum struct {
	INSTANCES ShowInstanceResourceInstancesCountRequestResourceType
}

func GetShowInstanceResourceInstancesCountRequestResourceTypeEnum() ShowInstanceResourceInstancesCountRequestResourceTypeEnum {
	return ShowInstanceResourceInstancesCountRequestResourceTypeEnum{
		INSTANCES: ShowInstanceResourceInstancesCountRequestResourceType{
			value: "instances",
		},
	}
}

func (c ShowInstanceResourceInstancesCountRequestResourceType) Value() string {
	return c.value
}

func (c ShowInstanceResourceInstancesCountRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceResourceInstancesCountRequestResourceType) UnmarshalJSON(b []byte) error {
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
