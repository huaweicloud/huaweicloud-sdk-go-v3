package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowInstanceExtendProductInfoRequest Request Object
type ShowInstanceExtendProductInfoRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 产品的类型。   - advanced：专享版
	Type ShowInstanceExtendProductInfoRequestType `json:"type"`

	// 消息引擎的类型。支持的类型为rabbitmq。
	Engine ShowInstanceExtendProductInfoRequestEngine `json:"engine"`
}

func (o ShowInstanceExtendProductInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceExtendProductInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceExtendProductInfoRequest", string(data)}, " ")
}

type ShowInstanceExtendProductInfoRequestType struct {
	value string
}

type ShowInstanceExtendProductInfoRequestTypeEnum struct {
	ADVANCED ShowInstanceExtendProductInfoRequestType
	PLATINUM ShowInstanceExtendProductInfoRequestType
	DEC      ShowInstanceExtendProductInfoRequestType
	EXP      ShowInstanceExtendProductInfoRequestType
}

func GetShowInstanceExtendProductInfoRequestTypeEnum() ShowInstanceExtendProductInfoRequestTypeEnum {
	return ShowInstanceExtendProductInfoRequestTypeEnum{
		ADVANCED: ShowInstanceExtendProductInfoRequestType{
			value: "advanced",
		},
		PLATINUM: ShowInstanceExtendProductInfoRequestType{
			value: "platinum",
		},
		DEC: ShowInstanceExtendProductInfoRequestType{
			value: "dec",
		},
		EXP: ShowInstanceExtendProductInfoRequestType{
			value: "exp",
		},
	}
}

func (c ShowInstanceExtendProductInfoRequestType) Value() string {
	return c.value
}

func (c ShowInstanceExtendProductInfoRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceExtendProductInfoRequestType) UnmarshalJSON(b []byte) error {
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

type ShowInstanceExtendProductInfoRequestEngine struct {
	value string
}

type ShowInstanceExtendProductInfoRequestEngineEnum struct {
	RABBITMQ ShowInstanceExtendProductInfoRequestEngine
}

func GetShowInstanceExtendProductInfoRequestEngineEnum() ShowInstanceExtendProductInfoRequestEngineEnum {
	return ShowInstanceExtendProductInfoRequestEngineEnum{
		RABBITMQ: ShowInstanceExtendProductInfoRequestEngine{
			value: "rabbitmq",
		},
	}
}

func (c ShowInstanceExtendProductInfoRequestEngine) Value() string {
	return c.value
}

func (c ShowInstanceExtendProductInfoRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceExtendProductInfoRequestEngine) UnmarshalJSON(b []byte) error {
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
