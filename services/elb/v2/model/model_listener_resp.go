package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 监听器响应体
type ListenerResp struct {
	// 监听器ID

	Id string `json:"id"`
	// 监听器所在的项目ID。

	TenantId string `json:"tenant_id"`
	// 监听器名称。

	Name string `json:"name"`
	// 监听器的描述信息

	Description string `json:"description"`
	// 监听器的管理状态。只支持设定为true，该字段的值无实际意义。

	AdminStateUp bool `json:"admin_state_up"`
	// 监听器绑定的负载均衡器ID的列表。

	Loadbalancers []ResourceList `json:"loadbalancers"`
	// 监听器的最大连接数。该字段为预留字段，暂未启用。默认为-1。

	ConnectionLimit int32 `json:"connection_limit"`
	// HTTP2功能的开启状态。该字段只有当监听器的协议是TERMINATED_HTTPS时生效。

	Http2Enable bool `json:"http2_enable"`
	// 监听器的监听协议

	Protocol ListenerRespProtocol `json:"protocol"`
	// 监听器的监听端口。

	ProtocolPort int32 `json:"protocol_port"`
	// 监听器的默认后端云服务器组ID。当请求没有匹配的转发策略时，转发到默认后端云服务器上处理。

	DefaultPoolId string `json:"default_pool_id"`
	// 监听器使用的服务器证书ID。

	DefaultTlsContainerRef string `json:"default_tls_container_ref"`
	// 监听器使用的CA证书ID。

	ClientCaTlsContainerRef string `json:"client_ca_tls_container_ref"`
	// 监听器使用的SNI证书（带域名的服务器证书）ID的列表。

	SniContainerRefs []string `json:"sni_container_refs"`
	// 监听器的标签。

	Tags []string `json:"tags"`
	// 监听器的创建时间。

	CreatedAt string `json:"created_at"`
	// 监听器的更新时间。

	UpdatedAt string `json:"updated_at"`

	InsertHeaders *InsertHeader `json:"insert_headers"`
	// 监听器所在的项目ID。

	ProjectId string `json:"project_id"`
	// 监听器使用的安全策略，仅对TERMINATED_HTTPS协议类型的监听器有效，且默认值为tls-1-0。  取值包括：tls-1-0, tls-1-1, tls-1-2, tls-1-2-strict多种安全策略

	TlsCiphersPolicy string `json:"tls_ciphers_policy"`
}

func (o ListenerResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListenerResp struct{}"
	}

	return strings.Join([]string{"ListenerResp", string(data)}, " ")
}

type ListenerRespProtocol struct {
	value string
}

type ListenerRespProtocolEnum struct {
	UDP              ListenerRespProtocol
	TCP              ListenerRespProtocol
	HTTP             ListenerRespProtocol
	TERMINATED_HTTPS ListenerRespProtocol
}

func GetListenerRespProtocolEnum() ListenerRespProtocolEnum {
	return ListenerRespProtocolEnum{
		UDP: ListenerRespProtocol{
			value: "UDP",
		},
		TCP: ListenerRespProtocol{
			value: "TCP",
		},
		HTTP: ListenerRespProtocol{
			value: "HTTP",
		},
		TERMINATED_HTTPS: ListenerRespProtocol{
			value: "TERMINATED_HTTPS",
		},
	}
}

func (c ListenerRespProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListenerRespProtocol) UnmarshalJSON(b []byte) error {
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
