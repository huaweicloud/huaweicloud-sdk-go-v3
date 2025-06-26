package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSubResourceInstancesCountRequest Request Object
type ShowSubResourceInstancesCountRequest struct {

	// 资源类型，支持的资源类型为：instances
	ResourceType ShowSubResourceInstancesCountRequestResourceType `json:"resource_type"`

	// 资源id
	ResourceId string `json:"resource_id"`

	// 子资源类型，支持的子资源类型为：namespaces
	SubResourceType ShowSubResourceInstancesCountRequestSubResourceType `json:"sub_resource_type"`

	Body *ListResourceInstancesRequestBody `json:"body,omitempty"`
}

func (o ShowSubResourceInstancesCountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubResourceInstancesCountRequest struct{}"
	}

	return strings.Join([]string{"ShowSubResourceInstancesCountRequest", string(data)}, " ")
}

type ShowSubResourceInstancesCountRequestResourceType struct {
	value string
}

type ShowSubResourceInstancesCountRequestResourceTypeEnum struct {
	INSTANCES ShowSubResourceInstancesCountRequestResourceType
}

func GetShowSubResourceInstancesCountRequestResourceTypeEnum() ShowSubResourceInstancesCountRequestResourceTypeEnum {
	return ShowSubResourceInstancesCountRequestResourceTypeEnum{
		INSTANCES: ShowSubResourceInstancesCountRequestResourceType{
			value: "instances",
		},
	}
}

func (c ShowSubResourceInstancesCountRequestResourceType) Value() string {
	return c.value
}

func (c ShowSubResourceInstancesCountRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSubResourceInstancesCountRequestResourceType) UnmarshalJSON(b []byte) error {
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

type ShowSubResourceInstancesCountRequestSubResourceType struct {
	value string
}

type ShowSubResourceInstancesCountRequestSubResourceTypeEnum struct {
	NAMESPACES ShowSubResourceInstancesCountRequestSubResourceType
}

func GetShowSubResourceInstancesCountRequestSubResourceTypeEnum() ShowSubResourceInstancesCountRequestSubResourceTypeEnum {
	return ShowSubResourceInstancesCountRequestSubResourceTypeEnum{
		NAMESPACES: ShowSubResourceInstancesCountRequestSubResourceType{
			value: "namespaces",
		},
	}
}

func (c ShowSubResourceInstancesCountRequestSubResourceType) Value() string {
	return c.value
}

func (c ShowSubResourceInstancesCountRequestSubResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSubResourceInstancesCountRequestSubResourceType) UnmarshalJSON(b []byte) error {
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
