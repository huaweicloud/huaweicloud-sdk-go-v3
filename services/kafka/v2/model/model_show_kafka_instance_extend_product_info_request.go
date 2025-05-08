package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowKafkaInstanceExtendProductInfoRequest Request Object
type ShowKafkaInstanceExtendProductInfoRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 产品的类型。 - advanced: 专享版
	Type ShowKafkaInstanceExtendProductInfoRequestType `json:"type"`
}

func (o ShowKafkaInstanceExtendProductInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaInstanceExtendProductInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowKafkaInstanceExtendProductInfoRequest", string(data)}, " ")
}

type ShowKafkaInstanceExtendProductInfoRequestType struct {
	value string
}

type ShowKafkaInstanceExtendProductInfoRequestTypeEnum struct {
	ADVANCED ShowKafkaInstanceExtendProductInfoRequestType
}

func GetShowKafkaInstanceExtendProductInfoRequestTypeEnum() ShowKafkaInstanceExtendProductInfoRequestTypeEnum {
	return ShowKafkaInstanceExtendProductInfoRequestTypeEnum{
		ADVANCED: ShowKafkaInstanceExtendProductInfoRequestType{
			value: "advanced",
		},
	}
}

func (c ShowKafkaInstanceExtendProductInfoRequestType) Value() string {
	return c.value
}

func (c ShowKafkaInstanceExtendProductInfoRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowKafkaInstanceExtendProductInfoRequestType) UnmarshalJSON(b []byte) error {
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
