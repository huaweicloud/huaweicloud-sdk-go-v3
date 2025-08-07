package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAssociatedResourcesResponse Response Object
type ShowAssociatedResourcesResponse struct {

	// 资源ID
	Id *string `json:"id,omitempty"`

	// 资源名称
	Name *string `json:"name,omitempty"`

	// 关联的资源类型
	Type *ShowAssociatedResourcesResponseType `json:"type,omitempty"`

	// 关联资源列表
	AssociatedResources *[]AssociatedResourceListResponse `json:"associated_resources,omitempty"`

	// 错误信息列表
	Errors         *[]ResourceErrorListResponse `json:"errors,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ShowAssociatedResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssociatedResourcesResponse struct{}"
	}

	return strings.Join([]string{"ShowAssociatedResourcesResponse", string(data)}, " ")
}

type ShowAssociatedResourcesResponseType struct {
	value string
}

type ShowAssociatedResourcesResponseTypeEnum struct {
	ECS ShowAssociatedResourcesResponseType
	EIP ShowAssociatedResourcesResponseType
}

func GetShowAssociatedResourcesResponseTypeEnum() ShowAssociatedResourcesResponseTypeEnum {
	return ShowAssociatedResourcesResponseTypeEnum{
		ECS: ShowAssociatedResourcesResponseType{
			value: "ecs",
		},
		EIP: ShowAssociatedResourcesResponseType{
			value: "eip",
		},
	}
}

func (c ShowAssociatedResourcesResponseType) Value() string {
	return c.value
}

func (c ShowAssociatedResourcesResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAssociatedResourcesResponseType) UnmarshalJSON(b []byte) error {
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
