package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 服务器配置
type CloudWafServer struct {
	// 对外协议

	FrontProtocol CloudWafServerFrontProtocol `json:"front_protocol"`
	// 源站协议

	BackProtocol CloudWafServerBackProtocol `json:"back_protocol"`
	// 源站地址

	Address string `json:"address"`
	// 源站端口

	Port int32 `json:"port"`
	// 源站地址为ipv4或ipv6

	Type CloudWafServerType `json:"type"`
}

func (o CloudWafServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudWafServer struct{}"
	}

	return strings.Join([]string{"CloudWafServer", string(data)}, " ")
}

type CloudWafServerFrontProtocol struct {
	value string
}

type CloudWafServerFrontProtocolEnum struct {
	HTTP  CloudWafServerFrontProtocol
	HTTPS CloudWafServerFrontProtocol
}

func GetCloudWafServerFrontProtocolEnum() CloudWafServerFrontProtocolEnum {
	return CloudWafServerFrontProtocolEnum{
		HTTP: CloudWafServerFrontProtocol{
			value: "HTTP",
		},
		HTTPS: CloudWafServerFrontProtocol{
			value: "HTTPS",
		},
	}
}

func (c CloudWafServerFrontProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CloudWafServerFrontProtocol) UnmarshalJSON(b []byte) error {
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

type CloudWafServerBackProtocol struct {
	value string
}

type CloudWafServerBackProtocolEnum struct {
	HTTP  CloudWafServerBackProtocol
	HTTPS CloudWafServerBackProtocol
}

func GetCloudWafServerBackProtocolEnum() CloudWafServerBackProtocolEnum {
	return CloudWafServerBackProtocolEnum{
		HTTP: CloudWafServerBackProtocol{
			value: "HTTP",
		},
		HTTPS: CloudWafServerBackProtocol{
			value: "HTTPS",
		},
	}
}

func (c CloudWafServerBackProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CloudWafServerBackProtocol) UnmarshalJSON(b []byte) error {
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

type CloudWafServerType struct {
	value string
}

type CloudWafServerTypeEnum struct {
	IPV4 CloudWafServerType
	IPV6 CloudWafServerType
}

func GetCloudWafServerTypeEnum() CloudWafServerTypeEnum {
	return CloudWafServerTypeEnum{
		IPV4: CloudWafServerType{
			value: "ipv4",
		},
		IPV6: CloudWafServerType{
			value: "ipv6",
		},
	}
}

func (c CloudWafServerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CloudWafServerType) UnmarshalJSON(b []byte) error {
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
