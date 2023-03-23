package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowEngineInstanceExtendProductInfoRequest struct {

	// 消息引擎的类型。支持的类型为rabbitmq。
	Engine string `json:"engine"`

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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
