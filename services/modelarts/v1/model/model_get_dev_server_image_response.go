package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GetDevServerImageResponse Response Object
type GetDevServerImageResponse struct {

	// **参数解释**：服务器镜像架构类型。 **取值范围**： - ARM - X86
	Arch *GetDevServerImageResponseArch `json:"arch,omitempty"`

	// **参数解释**：服务器镜像ID。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**：服务器镜像名称。表示服务器镜像的名称。 **约束限制**：不涉及。 **取值范围**：1 - 256字符 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：服务器类型。 **取值范围**： - BMS：裸金属服务器 - ECS：弹性云服务器 - HPS：超节点服务器
	ServerType *GetDevServerImageResponseServerType `json:"server_type,omitempty"`

	// **参数解释**：服务器镜像状态。 **取值范围**： - ACTIVE - INACTIVE
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GetDevServerImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetDevServerImageResponse struct{}"
	}

	return strings.Join([]string{"GetDevServerImageResponse", string(data)}, " ")
}

type GetDevServerImageResponseArch struct {
	value string
}

type GetDevServerImageResponseArchEnum struct {
	ARM GetDevServerImageResponseArch
	X86 GetDevServerImageResponseArch
}

func GetGetDevServerImageResponseArchEnum() GetDevServerImageResponseArchEnum {
	return GetDevServerImageResponseArchEnum{
		ARM: GetDevServerImageResponseArch{
			value: "ARM",
		},
		X86: GetDevServerImageResponseArch{
			value: "X86",
		},
	}
}

func (c GetDevServerImageResponseArch) Value() string {
	return c.value
}

func (c GetDevServerImageResponseArch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetDevServerImageResponseArch) UnmarshalJSON(b []byte) error {
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

type GetDevServerImageResponseServerType struct {
	value string
}

type GetDevServerImageResponseServerTypeEnum struct {
	BMS GetDevServerImageResponseServerType
	ECS GetDevServerImageResponseServerType
	HPS GetDevServerImageResponseServerType
}

func GetGetDevServerImageResponseServerTypeEnum() GetDevServerImageResponseServerTypeEnum {
	return GetDevServerImageResponseServerTypeEnum{
		BMS: GetDevServerImageResponseServerType{
			value: "BMS",
		},
		ECS: GetDevServerImageResponseServerType{
			value: "ECS",
		},
		HPS: GetDevServerImageResponseServerType{
			value: "HPS",
		},
	}
}

func (c GetDevServerImageResponseServerType) Value() string {
	return c.value
}

func (c GetDevServerImageResponseServerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetDevServerImageResponseServerType) UnmarshalJSON(b []byte) error {
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
