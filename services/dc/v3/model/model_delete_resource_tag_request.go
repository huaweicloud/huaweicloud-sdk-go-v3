package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteResourceTagRequest Request Object
type DeleteResourceTagRequest struct {

	// 标签key。
	Key string `json:"key"`

	// 资源实例ID
	ResourceId string `json:"resource_id"`

	// - 专线服务资源类型，包括dc-directconnect/dc-vgw/dc-vif - dc-directconnect: 专线物理连接 - dc-vgw： 虚拟网关 - dc-vif： 虚拟接口
	ResourceType DeleteResourceTagRequestResourceType `json:"resource_type"`
}

func (o DeleteResourceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteResourceTagRequest", string(data)}, " ")
}

type DeleteResourceTagRequestResourceType struct {
	value string
}

type DeleteResourceTagRequestResourceTypeEnum struct {
	DC_DIRECTCONNECT DeleteResourceTagRequestResourceType
	DC_VGW           DeleteResourceTagRequestResourceType
	DC_VIF           DeleteResourceTagRequestResourceType
	DC_LAG           DeleteResourceTagRequestResourceType
}

func GetDeleteResourceTagRequestResourceTypeEnum() DeleteResourceTagRequestResourceTypeEnum {
	return DeleteResourceTagRequestResourceTypeEnum{
		DC_DIRECTCONNECT: DeleteResourceTagRequestResourceType{
			value: "dc-directconnect",
		},
		DC_VGW: DeleteResourceTagRequestResourceType{
			value: "dc-vgw",
		},
		DC_VIF: DeleteResourceTagRequestResourceType{
			value: "dc-vif",
		},
		DC_LAG: DeleteResourceTagRequestResourceType{
			value: "dc-lag",
		},
	}
}

func (c DeleteResourceTagRequestResourceType) Value() string {
	return c.value
}

func (c DeleteResourceTagRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteResourceTagRequestResourceType) UnmarshalJSON(b []byte) error {
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
