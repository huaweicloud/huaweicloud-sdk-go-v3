package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type GatewayDto struct {

	// 网关类型
	ApigType *GatewayDtoApigType `json:"apig_type,omitempty"`

	// 网关实例id
	ApigInstanceId *string `json:"apig_instance_id,omitempty"`

	// 网关分组id
	GroupIdInApig *string `json:"group_id_in_apig,omitempty"`

	// roma网关集成应用id
	RomaAppId *string `json:"roma_app_id,omitempty"`
}

func (o GatewayDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GatewayDto struct{}"
	}

	return strings.Join([]string{"GatewayDto", string(data)}, " ")
}

type GatewayDtoApigType struct {
	value string
}

type GatewayDtoApigTypeEnum struct {
	APIG      GatewayDtoApigType
	APIGW     GatewayDtoApigType
	ROMA_APIC GatewayDtoApigType
}

func GetGatewayDtoApigTypeEnum() GatewayDtoApigTypeEnum {
	return GatewayDtoApigTypeEnum{
		APIG: GatewayDtoApigType{
			value: "APIG",
		},
		APIGW: GatewayDtoApigType{
			value: "APIGW",
		},
		ROMA_APIC: GatewayDtoApigType{
			value: "ROMA_APIC",
		},
	}
}

func (c GatewayDtoApigType) Value() string {
	return c.value
}

func (c GatewayDtoApigType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GatewayDtoApigType) UnmarshalJSON(b []byte) error {
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
