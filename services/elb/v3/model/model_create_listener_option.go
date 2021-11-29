package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建监听器的请求参数。
type CreateListenerOption struct {
	// 监听器的管理状态。固定为true。  不支持该字段，请勿使用。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 监听器默认的后端云服务器组ID。当请求没有匹配的转发策略时，转发到默认后端云服务器上处理。

	DefaultPoolId *string `json:"default_pool_id,omitempty"`
	// 监听器使用的CA证书ID。当且仅当type=client时，才会使用该字段对应的证书。

	ClientCaTlsContainerRef *string `json:"client_ca_tls_container_ref,omitempty"`
	// 监听器使用的服务器证书ID。  使用说明： - 当监听器协议为HTTPS时，该字段必传，且对应的证书的type必须是server类型。

	DefaultTlsContainerRef *string `json:"default_tls_container_ref,omitempty"`
	// 监听器的描述信息。

	Description *string `json:"description,omitempty"`
	// 客户端与LB之间的HTTPS请求的HTTP2功能的开启状态。开启后，可提升客户端与LB间的访问性能，但LB与后端服务器间仍采用HTTP1.X协议。 其他协议的监听器该字段无效，无论取值如何都不影响监听器正常运行。

	Http2Enable *bool `json:"http2_enable,omitempty"`

	InsertHeaders *ListenerInsertHeaders `json:"insert_headers,omitempty"`
	// 监听器所属的负载均衡器的ID列表。  使用说明： - 一个监听器只支持关联到一个LB。

	LoadbalancerId string `json:"loadbalancer_id"`
	// 监听器的名称。

	Name *string `json:"name,omitempty"`
	// 监听器所在的项目ID。

	ProjectId *string `json:"project_id,omitempty"`
	// 监听器的监听协议。支持TCP、UDP、HTTP、HTTPS、TERMINATED HTTPS。

	Protocol string `json:"protocol"`
	// 监听器的监听端口。

	ProtocolPort int32 `json:"protocol_port"`
	// 监听器使用的SNI证书（带域名的服务器证书）ID列表。  使用说明： - 列表对应的所有SNI证书的域名不允许存在重复。 - 列表对应的所有SNI证书的域名总数不超过30。

	SniContainerRefs *[]string `json:"sni_container_refs,omitempty"`
	// 标签列表

	Tags *[]Tag `json:"tags,omitempty"`
	// 监听器使用的安全策略，仅对HTTPS协议类型的监听器有效。 [取值：tls-1-0-inherit,tls-1-0, tls-1-1, tls-1-2,tls-1-2-strict，tls-1-2-fs，tls-1-0-with-1-3, tls-1-2-fs-with-1-3, hybrid-policy-1-0，默认：tls-1-0。](tag:hc,hws,hk,ocb,tlf,ctc,hcso,sbc,g42,tm,cmcc,hk-g42) [取值：tls-1-0, tls-1-1, tls-1-2, tls-1-2-strict，默认：tls-1-0。](tag:otc,otc_test,dt,dt_test) 使用说明： - 若同时设置了security_policy_id和tls_ciphers_policy，则仅security_policy_id生效。

	TlsCiphersPolicy *string `json:"tls_ciphers_policy,omitempty"`
	// 自定义安全策略的ID。仅关联LB为独享型时有效。  若同时设置了security_policy_id和tls_ciphers_policy，则仅security_policy_id生效。

	SecurityPolicyId *string `json:"security_policy_id,omitempty"`
	// 是否开启后端服务器的重试。取值：true 开启重试，false 不开启重试。默认：true。  使用说明： - 若关联是共享型LB，仅在protocol为HTTP、TERMINATED_HTTPS时才能传入该字段。 - 若关联是独享型LB，仅在protocol为HTTP、HTTPS时才能传入该字段。

	EnableMemberRetry *bool `json:"enable_member_retry,omitempty"`
	// 客户端连接空闲超时时间。在超过keepalive_timeout时长一直没有请求，负载均衡会暂时中断当前连接，直到一下次请求时重新建立新的连接。取值： - 若为TCP监听器，取值范围为（10-4000s）默认值为300s。 - 若为HTTP/HTTPS/TERMINATED_HTTPS监听器，取值范围为（0-4000s）默认值为60s。 UDP监听器不支持此字段。

	KeepaliveTimeout *int32 `json:"keepalive_timeout,omitempty"`
	// 等待客户端请求超时时间，包括两种情况： - 读取整个客户端请求头的超时时长：如果客户端未在超时时长内发送完整个请求头，则请求将被中断 - 两个连续body体的数据包到达LB的时间间隔，超出client_timeout将会断开连接。  取值范围为1-300s，默认值为60s。  使用说明：仅协议为HTTP/HTTPS的监听器支持该字段。

	ClientTimeout *int32 `json:"client_timeout,omitempty"`
	// 等待后端服务器响应超时时间。请求转发后端服务器后，在等待超时member_timeout时长没有响应，负载均衡将终止等待，并返回 HTTP504错误码。  取值：1-300s，默认为60s。  使用说明： - 仅支持协议为HTTP/HTTPS的监听器。

	MemberTimeout *int32 `json:"member_timeout,omitempty"`

	Ipgroup *CreateListenerIpGroupOption `json:"ipgroup,omitempty"`
	// 是否透传客户端IP地址。开启后客户端IP地址将透传到后端服务器。仅作用于共享型LB的TCP/UDP监听器。取值： - 共享型LB的TCP/UDP监听器可设置为true或false，不传默认为false。 - 共享型LB的HTTP/HTTPS监听器只支持设置为true，不传默认为true。 - 独享型负载均衡器所有协议的监听器只支持设置为true，不传默认为true。  使用说明： - 开启特性后，ELB和后端服务器之间直接使用真实的IP访问，需确保置服务器的安全组以及访问控制策略设置正确。 - 开启特性后，不支持同一台服务器既作为后端服务器又作为客户端的场景。 - 开启特性后，不支持变更后端服务器规格。

	TransparentClientIpEnable *bool `json:"transparent_client_ip_enable,omitempty"`
	// 是否开启高级转发策略功能。开启高级转发策略后，支持更灵活的转发策略和转发规则设置。取值：true开启，false不开启，默认false。 开启后支持如下场景：  - 转发策略的action字段支持指定为REDIRECT_TO_URL, FIXED_RESPONSE，即支持URL重定向和响应固定的内容给客户端。  - 转发策略支持指定priority、redirect_url_config、fixed_response_config字段。  - 转发规则rule的type可以指定METHOD, HEADER, QUERY_STRING, SOURCE_IP这几种取值。  - 转发规则rule的type为HOST_NAME时，转发规则rule的value支持通配符*。  - 转发规则支持指定conditions字段。

	EnhanceL7policyEnable *bool `json:"enhance_l7policy_enable,omitempty"`
}

func (o CreateListenerOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerOption struct{}"
	}

	return strings.Join([]string{"CreateListenerOption", string(data)}, " ")
}
