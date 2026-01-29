package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AlertConfig 订单配置开关项
type AlertConfig struct {

	// 关联对应的smn topic；当type=smn时，可以填写当前参数默认为空
	TopicUrn *string `json:"topic_urn,omitempty"`

	// 支持的订单告警类型，SMN为消息协议，MC为消息中心
	Type *AlertConfigType `json:"type,omitempty"`

	// 开启或关闭当前告警配置，true为开启，false为关闭
	Enable *bool `json:"enable,omitempty"`
}

func (o AlertConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertConfig struct{}"
	}

	return strings.Join([]string{"AlertConfig", string(data)}, " ")
}

type AlertConfigType struct {
	value string
}

type AlertConfigTypeEnum struct {
	SMN AlertConfigType
	MC  AlertConfigType
}

func GetAlertConfigTypeEnum() AlertConfigTypeEnum {
	return AlertConfigTypeEnum{
		SMN: AlertConfigType{
			value: "SMN",
		},
		MC: AlertConfigType{
			value: "MC",
		},
	}
}

func (c AlertConfigType) Value() string {
	return c.value
}

func (c AlertConfigType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlertConfigType) UnmarshalJSON(b []byte) error {
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
