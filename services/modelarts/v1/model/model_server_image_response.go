package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ServerImageResponse struct {

	// **参数解释**：服务器镜像架构类型。 **取值范围**： - ARM - X86
	Arch *ServerImageResponseArch `json:"arch,omitempty"`

	// **参数解释**：服务器镜像ID。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**：服务器镜像名称。表示服务器镜像的名称。 **约束限制**：不涉及。 **取值范围**：1 - 256字符 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：服务器类型。 **取值范围**： - BMS：裸金属服务器 - ECS：弹性云服务器 - HPS：超节点服务器
	ServerType *ServerImageResponseServerType `json:"server_type,omitempty"`

	// **参数解释**：服务器镜像状态。 **取值范围**： - ACTIVE - INACTIVE
	Status *string `json:"status,omitempty"`
}

func (o ServerImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerImageResponse struct{}"
	}

	return strings.Join([]string{"ServerImageResponse", string(data)}, " ")
}

type ServerImageResponseArch struct {
	value string
}

type ServerImageResponseArchEnum struct {
	ARM ServerImageResponseArch
	X86 ServerImageResponseArch
}

func GetServerImageResponseArchEnum() ServerImageResponseArchEnum {
	return ServerImageResponseArchEnum{
		ARM: ServerImageResponseArch{
			value: "ARM",
		},
		X86: ServerImageResponseArch{
			value: "X86",
		},
	}
}

func (c ServerImageResponseArch) Value() string {
	return c.value
}

func (c ServerImageResponseArch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServerImageResponseArch) UnmarshalJSON(b []byte) error {
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

type ServerImageResponseServerType struct {
	value string
}

type ServerImageResponseServerTypeEnum struct {
	BMS ServerImageResponseServerType
	ECS ServerImageResponseServerType
	HPS ServerImageResponseServerType
}

func GetServerImageResponseServerTypeEnum() ServerImageResponseServerTypeEnum {
	return ServerImageResponseServerTypeEnum{
		BMS: ServerImageResponseServerType{
			value: "BMS",
		},
		ECS: ServerImageResponseServerType{
			value: "ECS",
		},
		HPS: ServerImageResponseServerType{
			value: "HPS",
		},
	}
}

func (c ServerImageResponseServerType) Value() string {
	return c.value
}

func (c ServerImageResponseServerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServerImageResponseServerType) UnmarshalJSON(b []byte) error {
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
