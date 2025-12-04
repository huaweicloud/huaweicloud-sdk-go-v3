package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CloudWafServer 防护域名的源站服务器配置信息
type CloudWafServer struct {

	// **参数解释：** 客户端请求访问防护域名源站服务器的协议 **约束限制：** 不涉及 **取值范围：**  - HTTP: HTTP协议  - HTTPS: HTTPS协议  **默认取值：** 不涉及
	FrontProtocol CloudWafServerFrontProtocol `json:"front_protocol"`

	// **参数解释：** WAF转发客户端请求到防护域名源站服务器的协议 **约束限制：** 不涉及 **取值范围：**  - HTTP: HTTP协议  - HTTPS: HTTPS协议  **默认取值：** 不涉及
	BackProtocol CloudWafServerBackProtocol `json:"back_protocol"`

	// 源站权重，负载均衡算法将按该权重将请求分配给源站，默认值是1，云模式的冗余字段
	Weight *int32 `json:"weight,omitempty"`

	// 客户端访问的源站服务器的IP地址
	Address string `json:"address"`

	// WAF转发客户端请求到源站服务的业务端口
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

func (c CloudWafServerFrontProtocol) Value() string {
	return c.value
}

func (c CloudWafServerFrontProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CloudWafServerFrontProtocol) UnmarshalJSON(b []byte) error {
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

func (c CloudWafServerBackProtocol) Value() string {
	return c.value
}

func (c CloudWafServerBackProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CloudWafServerBackProtocol) UnmarshalJSON(b []byte) error {
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

func (c CloudWafServerType) Value() string {
	return c.value
}

func (c CloudWafServerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CloudWafServerType) UnmarshalJSON(b []byte) error {
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
