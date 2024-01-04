package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccessConfigurationPort 访问方式配置端口、协议、证书等信息。
type AccessConfigurationPort struct {

	// 监听端口。
	TargetPort *string `json:"target_port,omitempty"`

	// 访问端口。
	Port *string `json:"port,omitempty"`

	// 协议，负载均衡支持TCP，负载均衡与路由配置支持HTTP、HTTPS。
	Protocol *AccessConfigurationPortProtocol `json:"protocol,omitempty"`

	// 默认证书，访问方式配置为转发策略且协议为HTTPS时配置，未配置域名证书对时使用默认证书。
	DefaultCertificate *string `json:"default_certificate,omitempty"`

	// 证书。
	Certificate *string `json:"certificate,omitempty"`

	// 安全策略。
	Policy *AccessConfigurationPortPolicy `json:"policy,omitempty"`

	Path *[]AccessConfigurationHttpPath `json:"path,omitempty"`

	// 用户选择的elb的ID。
	ElbId *string `json:"elb_id,omitempty"`
}

func (o AccessConfigurationPort) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigurationPort struct{}"
	}

	return strings.Join([]string{"AccessConfigurationPort", string(data)}, " ")
}

type AccessConfigurationPortProtocol struct {
	value string
}

type AccessConfigurationPortProtocolEnum struct {
	TCP   AccessConfigurationPortProtocol
	HTTP  AccessConfigurationPortProtocol
	HTTPS AccessConfigurationPortProtocol
}

func GetAccessConfigurationPortProtocolEnum() AccessConfigurationPortProtocolEnum {
	return AccessConfigurationPortProtocolEnum{
		TCP: AccessConfigurationPortProtocol{
			value: "TCP",
		},
		HTTP: AccessConfigurationPortProtocol{
			value: "HTTP",
		},
		HTTPS: AccessConfigurationPortProtocol{
			value: "HTTPS",
		},
	}
}

func (c AccessConfigurationPortProtocol) Value() string {
	return c.value
}

func (c AccessConfigurationPortProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessConfigurationPortProtocol) UnmarshalJSON(b []byte) error {
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

type AccessConfigurationPortPolicy struct {
	value string
}

type AccessConfigurationPortPolicyEnum struct {
	TLS_1_2_STRICT AccessConfigurationPortPolicy
	TLS_1_2        AccessConfigurationPortPolicy
	TLS_1_1        AccessConfigurationPortPolicy
	TLS_1_0        AccessConfigurationPortPolicy
}

func GetAccessConfigurationPortPolicyEnum() AccessConfigurationPortPolicyEnum {
	return AccessConfigurationPortPolicyEnum{
		TLS_1_2_STRICT: AccessConfigurationPortPolicy{
			value: "tls-1-2-strict",
		},
		TLS_1_2: AccessConfigurationPortPolicy{
			value: "tls-1-2",
		},
		TLS_1_1: AccessConfigurationPortPolicy{
			value: "tls-1-1",
		},
		TLS_1_0: AccessConfigurationPortPolicy{
			value: "tls-1-0",
		},
	}
}

func (c AccessConfigurationPortPolicy) Value() string {
	return c.value
}

func (c AccessConfigurationPortPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessConfigurationPortPolicy) UnmarshalJSON(b []byte) error {
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
