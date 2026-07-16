package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ServerStartRequest struct {

	// **参数解释**：服务器架构信息。 **约束限制**：不涉及 **取值范围**： - -ARM - X86 **默认取值**：不涉及
	Arch *ServerStartRequestArch `json:"arch,omitempty"`

	// **参数解释**：服务器规格计费模式。 **约束限制**：不涉及。 **取值范围**： - [COMMON：同时支持包周期和按需](tag:hws,hws_hk) - POST_PAID：按需 - [PRE_PAID：包周期](tag:hws,hws_hk) **默认取值**：不涉及
	ChargingMode *ServerStartRequestChargingMode `json:"charging_mode,omitempty"`

	// **参数解释**：服务器类型。 **约束限制**：不涉及。 **取值范围**： - BMS：裸金属服务 - ECS：弹性云服务 - HPS：超节点服务 **默认取值**：不涉及
	ServerType *ServerStartRequestServerType `json:"server_type,omitempty"`
}

func (o ServerStartRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerStartRequest struct{}"
	}

	return strings.Join([]string{"ServerStartRequest", string(data)}, " ")
}

type ServerStartRequestArch struct {
	value string
}

type ServerStartRequestArchEnum struct {
	ARM ServerStartRequestArch
	X86 ServerStartRequestArch
}

func GetServerStartRequestArchEnum() ServerStartRequestArchEnum {
	return ServerStartRequestArchEnum{
		ARM: ServerStartRequestArch{
			value: "ARM",
		},
		X86: ServerStartRequestArch{
			value: "X86",
		},
	}
}

func (c ServerStartRequestArch) Value() string {
	return c.value
}

func (c ServerStartRequestArch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServerStartRequestArch) UnmarshalJSON(b []byte) error {
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

type ServerStartRequestChargingMode struct {
	value string
}

type ServerStartRequestChargingModeEnum struct {
	COMMON    ServerStartRequestChargingMode
	POST_PAID ServerStartRequestChargingMode
	PRE_PAID  ServerStartRequestChargingMode
}

func GetServerStartRequestChargingModeEnum() ServerStartRequestChargingModeEnum {
	return ServerStartRequestChargingModeEnum{
		COMMON: ServerStartRequestChargingMode{
			value: "COMMON",
		},
		POST_PAID: ServerStartRequestChargingMode{
			value: "POST_PAID",
		},
		PRE_PAID: ServerStartRequestChargingMode{
			value: "PRE_PAID",
		},
	}
}

func (c ServerStartRequestChargingMode) Value() string {
	return c.value
}

func (c ServerStartRequestChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServerStartRequestChargingMode) UnmarshalJSON(b []byte) error {
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

type ServerStartRequestServerType struct {
	value string
}

type ServerStartRequestServerTypeEnum struct {
	BMS ServerStartRequestServerType
	ECS ServerStartRequestServerType
	HPS ServerStartRequestServerType
}

func GetServerStartRequestServerTypeEnum() ServerStartRequestServerTypeEnum {
	return ServerStartRequestServerTypeEnum{
		BMS: ServerStartRequestServerType{
			value: "BMS",
		},
		ECS: ServerStartRequestServerType{
			value: "ECS",
		},
		HPS: ServerStartRequestServerType{
			value: "HPS",
		},
	}
}

func (c ServerStartRequestServerType) Value() string {
	return c.value
}

func (c ServerStartRequestServerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServerStartRequestServerType) UnmarshalJSON(b []byte) error {
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
