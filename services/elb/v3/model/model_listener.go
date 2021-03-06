package model

import (
	"encoding/json"

	"strings"
)

// 创建监听器的响应体
type Listener struct {
	// 监听器的管理状态。只支持设定为true，该字段的值无实际意义。

	AdminStateUp bool `json:"admin_state_up"`
	// 监听器使用的CA证书ID。

	ClientCaTlsContainerRef string `json:"client_ca_tls_container_ref"`
	// 监听器的最大连接数。该字段为预留字段，暂未启用。默认为-1。

	ConnectionLimit int32 `json:"connection_limit"`
	// 监听器的创建时间。

	CreatedAt string `json:"created_at"`
	// 监听器的默认后端云服务器组ID。当请求没有匹配的转发策略时，转发到默认后端云服务器上处理。

	DefaultPoolId string `json:"default_pool_id"`
	// 监听器使用的服务器证书ID。

	DefaultTlsContainerRef string `json:"default_tls_container_ref"`
	// 监听器的描述信息

	Description string `json:"description"`
	// HTTP2功能的开启状态。该字段只有当监听器的协议是TERMINATED_HTTPS时生效。

	Http2Enable bool `json:"http2_enable"`
	// 监听器ID

	Id string `json:"id"`

	InsertHeaders *ListenerInsertHeaders `json:"insert_headers"`
	// 监听器绑定的负载均衡器ID的列表。

	Loadbalancers []LoadBalancerRef `json:"loadbalancers"`
	// 监听器名称

	Name string `json:"name"`
	// 监听器所在的项目ID。

	ProjectId string `json:"project_id"`
	// 监听器的监听协议

	Protocol string `json:"protocol"`
	// 监听器的监听端口。

	ProtocolPort int32 `json:"protocol_port"`
	// 监听器使用的SNI证书（带域名的服务器证书）ID的列表。

	SniContainerRefs []string `json:"sni_container_refs"`
	// 标签列表

	Tags []Tag `json:"tags"`
	// 监听器的更新时间。

	UpdatedAt string `json:"updated_at"`
	// 监听器使用的安全策略，仅对TERMINATED_HTTPS协议类型的监听器有效，且默认值为tls-1-0。 取值包括：tls-1-0-inherit,tls-1-0, tls-1-1, tls-1-2, tls-1-2-strict，tls-1-2-fs六种安全策略

	TlsCiphersPolicy string `json:"tls_ciphers_policy"`
	// 是否关闭后端服务器的重试。 该字段只在protocol为HTTP、HTTPS时有意义。

	EnableMemberRetry bool `json:"enable_member_retry"`
	// TCP监听器配置空闲超时时间，取值范围为（10-900s）默认值为300s，TCP监听器配置空闲超时时间，取值范围为（10-900s）默认值为300s，HTTP/TERMINATED_HTTPS监听器为客户端连接空闲超时时间，取值范围为（1-300s）默认值为15s。 UDP此字段无意义

	KeepaliveTimeout *int32 `json:"keepalive_timeout,omitempty"`
	// 等待客户端请求超时时间，协议为HTTP， TERMINATED_HTTPS的监听器才有意义。取值范围 1-60

	ClientTimeout *int32 `json:"client_timeout,omitempty"`
	// 等待后端服务器请求超时时间，协议为HTTP， TERMINATED_HTTPS时才有意义。取值范围 1-300

	MemberTimeout *int32 `json:"member_timeout,omitempty"`

	Ipgroup *ListenerIpGroup `json:"ipgroup"`
}

func (o Listener) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Listener struct{}"
	}

	return strings.Join([]string{"Listener", string(data)}, " ")
}
