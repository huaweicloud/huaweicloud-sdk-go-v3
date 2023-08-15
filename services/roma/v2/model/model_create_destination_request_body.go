package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateDestinationRequestBody struct {

	// 操作类型，枚举值:0-目标端为本ROMA实例内MQS，； 7-目标端为设备
	DestinationType CreateDestinationRequestBodyDestinationType `json:"destination_type"`

	// 应用ID，目标端为0时需明确对方的APP_ID
	AppId *string `json:"app_id,omitempty"`

	// 目标数据源名称
	DestinationName *string `json:"destination_name,omitempty"`

	// 目标数据源主题，从MQS服务中获取已有topic
	Topic string `json:"topic"`

	// 目标端数据源服务，连接地址
	Server *string `json:"server,omitempty"`

	// 目标端数据源token
	Token *string `json:"token,omitempty"`

	// 目标数据源标签
	Tag *string `json:"tag,omitempty"`

	// 目标端数据源MQS的SASL字段是否需要支持SSL加密
	MqsSaslSsl *bool `json:"mqs_sasl_ssl,omitempty"`

	// 目标数据源用户名
	UserName *string `json:"user_name,omitempty"`

	// 目标数据源密码
	Password *string `json:"password,omitempty"`
}

func (o CreateDestinationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDestinationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDestinationRequestBody", string(data)}, " ")
}

type CreateDestinationRequestBodyDestinationType struct {
	value int32
}

type CreateDestinationRequestBodyDestinationTypeEnum struct {
	E_0 CreateDestinationRequestBodyDestinationType
	E_7 CreateDestinationRequestBodyDestinationType
}

func GetCreateDestinationRequestBodyDestinationTypeEnum() CreateDestinationRequestBodyDestinationTypeEnum {
	return CreateDestinationRequestBodyDestinationTypeEnum{
		E_0: CreateDestinationRequestBodyDestinationType{
			value: 0,
		}, E_7: CreateDestinationRequestBodyDestinationType{
			value: 7,
		},
	}
}

func (c CreateDestinationRequestBodyDestinationType) Value() int32 {
	return c.value
}

func (c CreateDestinationRequestBodyDestinationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDestinationRequestBodyDestinationType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
