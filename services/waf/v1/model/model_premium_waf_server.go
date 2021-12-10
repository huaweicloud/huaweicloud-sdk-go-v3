package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 服务器配置
type PremiumWafServer struct {
	// 对外协议

	FrontProtocol PremiumWafServerFrontProtocol `json:"front_protocol"`
	// 源站协议

	BackProtocol PremiumWafServerBackProtocol `json:"back_protocol"`
	// 源站地址

	Address string `json:"address"`
	// 源站端口

	Port int32 `json:"port"`
	// 源站地址为ipv4或ipv6

	Type PremiumWafServerType `json:"type"`
	// VPC id,通过以下步骤获取VPC id： \\n 1.找到独享引擎所在的虚拟私有云名称，VPC\\子网这一列就是VPC的名称：登录WAF的控制台->单击系统管理->独享引擎->VPC\\子网 \\n 2.登录虚拟私有云 VPC控制台->虚拟私有云->单击虚拟私有云的名称->基本信息的ID

	VpcId string `json:"vpc_id"`
}

func (o PremiumWafServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PremiumWafServer struct{}"
	}

	return strings.Join([]string{"PremiumWafServer", string(data)}, " ")
}

type PremiumWafServerFrontProtocol struct {
	value string
}

type PremiumWafServerFrontProtocolEnum struct {
	HTTP  PremiumWafServerFrontProtocol
	HTTPS PremiumWafServerFrontProtocol
}

func GetPremiumWafServerFrontProtocolEnum() PremiumWafServerFrontProtocolEnum {
	return PremiumWafServerFrontProtocolEnum{
		HTTP: PremiumWafServerFrontProtocol{
			value: "HTTP",
		},
		HTTPS: PremiumWafServerFrontProtocol{
			value: "HTTPS",
		},
	}
}

func (c PremiumWafServerFrontProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PremiumWafServerFrontProtocol) UnmarshalJSON(b []byte) error {
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

type PremiumWafServerBackProtocol struct {
	value string
}

type PremiumWafServerBackProtocolEnum struct {
	HTTP  PremiumWafServerBackProtocol
	HTTPS PremiumWafServerBackProtocol
}

func GetPremiumWafServerBackProtocolEnum() PremiumWafServerBackProtocolEnum {
	return PremiumWafServerBackProtocolEnum{
		HTTP: PremiumWafServerBackProtocol{
			value: "HTTP",
		},
		HTTPS: PremiumWafServerBackProtocol{
			value: "HTTPS",
		},
	}
}

func (c PremiumWafServerBackProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PremiumWafServerBackProtocol) UnmarshalJSON(b []byte) error {
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

type PremiumWafServerType struct {
	value string
}

type PremiumWafServerTypeEnum struct {
	IPV4 PremiumWafServerType
	IPV6 PremiumWafServerType
}

func GetPremiumWafServerTypeEnum() PremiumWafServerTypeEnum {
	return PremiumWafServerTypeEnum{
		IPV4: PremiumWafServerType{
			value: "ipv4",
		},
		IPV6: PremiumWafServerType{
			value: "ipv6",
		},
	}
}

func (c PremiumWafServerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PremiumWafServerType) UnmarshalJSON(b []byte) error {
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
