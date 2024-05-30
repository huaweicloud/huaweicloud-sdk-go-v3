package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// WafServer 防护域名的源站服务器配置信息
type WafServer struct {

	// 客户端请求访问防护域名源站服务器的协议
	FrontProtocol *WafServerFrontProtocol `json:"front_protocol,omitempty"`

	// WAF转发客户端请求到防护域名源站服务器的协议
	BackProtocol *WafServerBackProtocol `json:"back_protocol,omitempty"`

	// 源站权重，负载均衡算法将按该权重将请求分配给源站，默认值是1，云模式的冗余字段
	Weight *int32 `json:"weight,omitempty"`

	// 客户端访问的源站服务器的IP地址
	Address *string `json:"address,omitempty"`

	// WAF转发客户端请求到源站服务的业务端口
	Port *int32 `json:"port,omitempty"`

	// 源站地址为ipv4或ipv6
	Type *WafServerType `json:"type,omitempty"`

	// VPC id,通过以下步骤获取VPC id：   - 1.找到独享引擎所在的虚拟私有云名称，VPC\\子网这一列就是VPC的名称：登录WAF的控制台->单击系统管理->独享引擎->VPC\\子网   - 2.登录虚拟私有云 VPC控制台->虚拟私有云->单击虚拟私有云的名称->基本信息的ID
	VpcId *string `json:"vpc_id,omitempty"`
}

func (o WafServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WafServer struct{}"
	}

	return strings.Join([]string{"WafServer", string(data)}, " ")
}

type WafServerFrontProtocol struct {
	value string
}

type WafServerFrontProtocolEnum struct {
	HTTP  WafServerFrontProtocol
	HTTPS WafServerFrontProtocol
}

func GetWafServerFrontProtocolEnum() WafServerFrontProtocolEnum {
	return WafServerFrontProtocolEnum{
		HTTP: WafServerFrontProtocol{
			value: "HTTP",
		},
		HTTPS: WafServerFrontProtocol{
			value: "HTTPS",
		},
	}
}

func (c WafServerFrontProtocol) Value() string {
	return c.value
}

func (c WafServerFrontProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WafServerFrontProtocol) UnmarshalJSON(b []byte) error {
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

type WafServerBackProtocol struct {
	value string
}

type WafServerBackProtocolEnum struct {
	HTTP  WafServerBackProtocol
	HTTPS WafServerBackProtocol
}

func GetWafServerBackProtocolEnum() WafServerBackProtocolEnum {
	return WafServerBackProtocolEnum{
		HTTP: WafServerBackProtocol{
			value: "HTTP",
		},
		HTTPS: WafServerBackProtocol{
			value: "HTTPS",
		},
	}
}

func (c WafServerBackProtocol) Value() string {
	return c.value
}

func (c WafServerBackProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WafServerBackProtocol) UnmarshalJSON(b []byte) error {
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

type WafServerType struct {
	value string
}

type WafServerTypeEnum struct {
	IPV4 WafServerType
	IPV6 WafServerType
}

func GetWafServerTypeEnum() WafServerTypeEnum {
	return WafServerTypeEnum{
		IPV4: WafServerType{
			value: "ipv4",
		},
		IPV6: WafServerType{
			value: "ipv6",
		},
	}
}

func (c WafServerType) Value() string {
	return c.value
}

func (c WafServerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WafServerType) UnmarshalJSON(b []byte) error {
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
