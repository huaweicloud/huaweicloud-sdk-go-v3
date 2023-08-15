package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowEngineInstanceExtendProductInfoRequest Request Object
type ShowEngineInstanceExtendProductInfoRequest struct {

	// 消息引擎的类型。支持的类型为rabbitmq。
	Engine ShowEngineInstanceExtendProductInfoRequestEngine `json:"engine"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 产品的类型。   - advanced：专享版   - platinum：铂金版   - dec：专属云版   - exp：体验版
	Type *ShowEngineInstanceExtendProductInfoRequestType `json:"type,omitempty"`
}

func (o ShowEngineInstanceExtendProductInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEngineInstanceExtendProductInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowEngineInstanceExtendProductInfoRequest", string(data)}, " ")
}

type ShowEngineInstanceExtendProductInfoRequestEngine struct {
	value string
}

type ShowEngineInstanceExtendProductInfoRequestEngineEnum struct {
	RABBITMQ ShowEngineInstanceExtendProductInfoRequestEngine
}

func GetShowEngineInstanceExtendProductInfoRequestEngineEnum() ShowEngineInstanceExtendProductInfoRequestEngineEnum {
	return ShowEngineInstanceExtendProductInfoRequestEngineEnum{
		RABBITMQ: ShowEngineInstanceExtendProductInfoRequestEngine{
			value: "rabbitmq",
		},
	}
}

func (c ShowEngineInstanceExtendProductInfoRequestEngine) Value() string {
	return c.value
}

func (c ShowEngineInstanceExtendProductInfoRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEngineInstanceExtendProductInfoRequestEngine) UnmarshalJSON(b []byte) error {
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

type ShowEngineInstanceExtendProductInfoRequestType struct {
	value string
}

type ShowEngineInstanceExtendProductInfoRequestTypeEnum struct {
	ADVANCED ShowEngineInstanceExtendProductInfoRequestType
	PLATINUM ShowEngineInstanceExtendProductInfoRequestType
	DEC      ShowEngineInstanceExtendProductInfoRequestType
	EXP      ShowEngineInstanceExtendProductInfoRequestType
}

func GetShowEngineInstanceExtendProductInfoRequestTypeEnum() ShowEngineInstanceExtendProductInfoRequestTypeEnum {
	return ShowEngineInstanceExtendProductInfoRequestTypeEnum{
		ADVANCED: ShowEngineInstanceExtendProductInfoRequestType{
			value: "advanced",
		},
		PLATINUM: ShowEngineInstanceExtendProductInfoRequestType{
			value: "platinum",
		},
		DEC: ShowEngineInstanceExtendProductInfoRequestType{
			value: "dec",
		},
		EXP: ShowEngineInstanceExtendProductInfoRequestType{
			value: "exp",
		},
	}
}

func (c ShowEngineInstanceExtendProductInfoRequestType) Value() string {
	return c.value
}

func (c ShowEngineInstanceExtendProductInfoRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEngineInstanceExtendProductInfoRequestType) UnmarshalJSON(b []byte) error {
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
