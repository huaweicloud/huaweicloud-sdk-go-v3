package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Destination struct {
	// 目标数据源ID

	DestinationId *int32 `json:"destination_id,omitempty"`
	// 操作类型，枚举值:0-目标端为本ROMA实例内MQS； 7-目标端为设备

	DestinationType *DestinationDestinationType `json:"destination_type,omitempty"`
	// 应用ID，目标端为0时需明确对方的APP_ID

	AppId *string `json:"app_id,omitempty"`
	// 目标数据源名称

	DestinationName *string `json:"destination_name,omitempty"`
	// 目标数据源主题

	Topic *string `json:"topic,omitempty"`
	// 目标端数据源服务

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

func (o Destination) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Destination struct{}"
	}

	return strings.Join([]string{"Destination", string(data)}, " ")
}

type DestinationDestinationType struct {
	value int32
}

type DestinationDestinationTypeEnum struct {
	E_0 DestinationDestinationType
	E_7 DestinationDestinationType
}

func GetDestinationDestinationTypeEnum() DestinationDestinationTypeEnum {
	return DestinationDestinationTypeEnum{
		E_0: DestinationDestinationType{
			value: 0,
		}, E_7: DestinationDestinationType{
			value: 7,
		},
	}
}

func (c DestinationDestinationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DestinationDestinationType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
