package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 服务器配置
type UpdateCloudWafServer struct {
	// 对外协议

	FrontProtocol *UpdateCloudWafServerFrontProtocol `json:"front_protocol,omitempty"`
	// 源站协议

	BackProtocol *UpdateCloudWafServerBackProtocol `json:"back_protocol,omitempty"`
	// 源站地址

	Address *string `json:"address,omitempty"`
	// 源站端口

	Port *int32 `json:"port,omitempty"`
	// 源站地址为ipv4或ipv6

	Type *UpdateCloudWafServerType `json:"type,omitempty"`
}

func (o UpdateCloudWafServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCloudWafServer struct{}"
	}

	return strings.Join([]string{"UpdateCloudWafServer", string(data)}, " ")
}

type UpdateCloudWafServerFrontProtocol struct {
	value string
}

type UpdateCloudWafServerFrontProtocolEnum struct {
	HTTP  UpdateCloudWafServerFrontProtocol
	HTTPS UpdateCloudWafServerFrontProtocol
}

func GetUpdateCloudWafServerFrontProtocolEnum() UpdateCloudWafServerFrontProtocolEnum {
	return UpdateCloudWafServerFrontProtocolEnum{
		HTTP: UpdateCloudWafServerFrontProtocol{
			value: "HTTP",
		},
		HTTPS: UpdateCloudWafServerFrontProtocol{
			value: "HTTPS",
		},
	}
}

func (c UpdateCloudWafServerFrontProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateCloudWafServerFrontProtocol) UnmarshalJSON(b []byte) error {
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

type UpdateCloudWafServerBackProtocol struct {
	value string
}

type UpdateCloudWafServerBackProtocolEnum struct {
	HTTP  UpdateCloudWafServerBackProtocol
	HTTPS UpdateCloudWafServerBackProtocol
}

func GetUpdateCloudWafServerBackProtocolEnum() UpdateCloudWafServerBackProtocolEnum {
	return UpdateCloudWafServerBackProtocolEnum{
		HTTP: UpdateCloudWafServerBackProtocol{
			value: "HTTP",
		},
		HTTPS: UpdateCloudWafServerBackProtocol{
			value: "HTTPS",
		},
	}
}

func (c UpdateCloudWafServerBackProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateCloudWafServerBackProtocol) UnmarshalJSON(b []byte) error {
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

type UpdateCloudWafServerType struct {
	value string
}

type UpdateCloudWafServerTypeEnum struct {
	IPV4 UpdateCloudWafServerType
	IPV6 UpdateCloudWafServerType
}

func GetUpdateCloudWafServerTypeEnum() UpdateCloudWafServerTypeEnum {
	return UpdateCloudWafServerTypeEnum{
		IPV4: UpdateCloudWafServerType{
			value: "ipv4",
		},
		IPV6: UpdateCloudWafServerType{
			value: "ipv6",
		},
	}
}

func (c UpdateCloudWafServerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateCloudWafServerType) UnmarshalJSON(b []byte) error {
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
