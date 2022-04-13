package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 创建监听器的结构体
type CreateListenerReq struct {
	// 监听器关联的负载均衡器 ID

	LoadbalancerId string `json:"loadbalancer_id"`
	// 监听器的监听协议

	Protocol CreateListenerReqProtocol `json:"protocol"`
	// 监听器的监听端口。如果监听协议为UDP，端口号不支持4789。

	ProtocolPort int32 `json:"protocol_port"`
	// 监听器所在的项目ID。

	TenantId *string `json:"tenant_id,omitempty"`
	// 监听器名称。

	Name *string `json:"name,omitempty"`
	// 监听器的描述信息

	Description *string `json:"description,omitempty"`
	// 监听器器的管理状态。只支持设定为true，该字段的值无实际意义。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 监听器的最大连接数。该字段为预留字段，暂未启用。默认为-1。

	ConnectionLimit *int32 `json:"connection_limit,omitempty"`
	// HTTP2功能的开启状态。该字段只有当监听器的协议是TERMINATED_HTTPS时才有意义。

	Http2Enable *bool `json:"http2_enable,omitempty"`
	// 监听器的默认后端云服务器组ID。当请求没有匹配的转发策略时，转发到默认后端云服务器上处理。当该字段为null时，表示监听器无默认的后端云服务器组。

	DefaultPoolId *string `json:"default_pool_id,omitempty"`
	// 监听器使用的服务器证书ID。当protocol参数为TERMINATED_HTTPS时，为必选字段

	DefaultTlsContainerRef *string `json:"default_tls_container_ref,omitempty"`
	// 监听器使用的CA证书ID。

	ClientCaTlsContainerRef *string `json:"client_ca_tls_container_ref,omitempty"`
	// 监听器使用的SNI证书（带域名的服务器证书）ID的列表。 该字段不为空列表时，SNI特性开启。该字段为空列表时，SNI特性关闭。

	SniContainerRefs *[]string `json:"sni_container_refs,omitempty"`

	InsertHeaders *InsertHeader `json:"insert_headers,omitempty"`
	// 监听器使用的安全策略，仅对TERMINATED_HTTPS协议类型的监听器有效，且默认值为tls-1-0。  取值包括：tls-1-0, tls-1-1, tls-1-2, tls-1-2-strict多种安全策略。

	TlsCiphersPolicy *CreateListenerReqTlsCiphersPolicy `json:"tls_ciphers_policy,omitempty"`
}

func (o CreateListenerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerReq struct{}"
	}

	return strings.Join([]string{"CreateListenerReq", string(data)}, " ")
}

type CreateListenerReqProtocol struct {
	value string
}

type CreateListenerReqProtocolEnum struct {
	UDP              CreateListenerReqProtocol
	TCP              CreateListenerReqProtocol
	HTTP             CreateListenerReqProtocol
	TERMINATED_HTTPS CreateListenerReqProtocol
}

func GetCreateListenerReqProtocolEnum() CreateListenerReqProtocolEnum {
	return CreateListenerReqProtocolEnum{
		UDP: CreateListenerReqProtocol{
			value: "UDP",
		},
		TCP: CreateListenerReqProtocol{
			value: "TCP",
		},
		HTTP: CreateListenerReqProtocol{
			value: "HTTP",
		},
		TERMINATED_HTTPS: CreateListenerReqProtocol{
			value: "TERMINATED_HTTPS",
		},
	}
}

func (c CreateListenerReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateListenerReqProtocol) UnmarshalJSON(b []byte) error {
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

type CreateListenerReqTlsCiphersPolicy struct {
	value string
}

type CreateListenerReqTlsCiphersPolicyEnum struct {
	TLS_1_0        CreateListenerReqTlsCiphersPolicy
	TLS_1_1        CreateListenerReqTlsCiphersPolicy
	TLS_1_2        CreateListenerReqTlsCiphersPolicy
	TLS_1_2_STRICT CreateListenerReqTlsCiphersPolicy
}

func GetCreateListenerReqTlsCiphersPolicyEnum() CreateListenerReqTlsCiphersPolicyEnum {
	return CreateListenerReqTlsCiphersPolicyEnum{
		TLS_1_0: CreateListenerReqTlsCiphersPolicy{
			value: "tls-1-0",
		},
		TLS_1_1: CreateListenerReqTlsCiphersPolicy{
			value: " tls-1-1",
		},
		TLS_1_2: CreateListenerReqTlsCiphersPolicy{
			value: " tls-1-2",
		},
		TLS_1_2_STRICT: CreateListenerReqTlsCiphersPolicy{
			value: " tls-1-2-strict",
		},
	}
}

func (c CreateListenerReqTlsCiphersPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateListenerReqTlsCiphersPolicy) UnmarshalJSON(b []byte) error {
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
