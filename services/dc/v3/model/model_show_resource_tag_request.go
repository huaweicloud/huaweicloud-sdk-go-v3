package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowResourceTagRequest Request Object
type ShowResourceTagRequest struct {

	// - 专线服务资源类型，包括dc-directconnect/dc-vgw/dc-vif - dc-directconnect: 专线物理连接 - dc-vgw： 虚拟网关 - dc-vif： 虚拟接口
	ResourceType ShowResourceTagRequestResourceType `json:"resource_type"`

	// 资源实例ID
	ResourceId string `json:"resource_id"`
}

func (o ShowResourceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceTagRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceTagRequest", string(data)}, " ")
}

type ShowResourceTagRequestResourceType struct {
	value string
}

type ShowResourceTagRequestResourceTypeEnum struct {
	DC_DIRECTCONNECT ShowResourceTagRequestResourceType
	DC_VGW           ShowResourceTagRequestResourceType
	DC_VIF           ShowResourceTagRequestResourceType
	DC_LAG           ShowResourceTagRequestResourceType
}

func GetShowResourceTagRequestResourceTypeEnum() ShowResourceTagRequestResourceTypeEnum {
	return ShowResourceTagRequestResourceTypeEnum{
		DC_DIRECTCONNECT: ShowResourceTagRequestResourceType{
			value: "dc-directconnect",
		},
		DC_VGW: ShowResourceTagRequestResourceType{
			value: "dc-vgw",
		},
		DC_VIF: ShowResourceTagRequestResourceType{
			value: "dc-vif",
		},
		DC_LAG: ShowResourceTagRequestResourceType{
			value: "dc-lag",
		},
	}
}

func (c ShowResourceTagRequestResourceType) Value() string {
	return c.value
}

func (c ShowResourceTagRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowResourceTagRequestResourceType) UnmarshalJSON(b []byte) error {
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
