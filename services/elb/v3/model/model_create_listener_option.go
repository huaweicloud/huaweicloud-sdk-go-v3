package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateListenerOption 创建监听器的请求参数。
type CreateListenerOption struct {

	// 参数解释：监听器的管理状态。  约束限制：只支持设置为true。  [不支持该字段，请勿使用。](tag:dt,dt_test,hcso_dt)
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 参数解释：监听器默认的后端云服务器组ID。当请求没有匹配的转发策略时，转发到默认后端云服务器上处理。
	DefaultPoolId *string `json:"default_pool_id,omitempty"`

	// 参数解释：监听器使用的CA证书ID。  约束限制： - 当且仅当type=client时，才会使用该字段对应的证书。 - 监听器协议为QUIC时不支持该字段。  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt,dt,dt_test)
	ClientCaTlsContainerRef *string `json:"client_ca_tls_container_ref,omitempty"`

	// 参数解释：监听器使用的服务器证书ID。  约束限制：当监听器协议为HTTPS、TLS或QUIC时，该字段必传，且对应的证书的type必须是server类型。
	DefaultTlsContainerRef *string `json:"default_tls_container_ref,omitempty"`

	// 参数解释：监听器的描述信息。
	Description *string `json:"description,omitempty"`

	// 参数解释：客户端与LB之间的HTTPS请求的HTTP2功能的开启状态。 开启后，可提升客户端与LB间的访问性能，但LB与后端服务器间仍采用HTTP1.X协议。  约束限制： - 仅HTTPS协议监听器有效。 - QUIC监听器不能设置该字段，固定返回为true。 - 其他协议的监听器可设置该字段但无效，无论取值如何都不影响监听器正常运行。  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt,dt,dt_test)
	Http2Enable *bool `json:"http2_enable,omitempty"`

	InsertHeaders *ListenerInsertHeaders `json:"insert_headers,omitempty"`

	// 参数解释：监听器所属的负载均衡器的ID列表。  约束限制： 一个监听器只支持关联到一个LB。
	LoadbalancerId string `json:"loadbalancer_id"`

	// 参数解释：监听器的名称。  注意：若名称为空，则在控制台的监听器列表无法选择并查看监听器详情。
	Name *string `json:"name,omitempty"`

	// 参数解释：监听器所在的项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 参数解释：监听器的监听协议。  [ 约束限制： - 共享型LB上的HTTPS监听器只支持设置为TERMINATED_HTTPS。 传入HTTPS将会自动转为TERMINATED_HTTPS。 - 独享型LB上的HTTPS监听器只支持设置为HTTPS，传入TERMINATED_HTTPS将会自动转为HTTPS。  取值范围：TCP、UDP、HTTP、HTTPS、TERMINATED_HTTPS、QUIC、TLS。 ](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs,dt)  [取值范围：TCP、UDP、HTTP、HTTPS。](tag:hcso_dt) [取值范围：TCP、UDP、IP、HTTP、HTTPS。IP为网关型LB上的监听器独有的协议。](tag:hws_eu)  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt,dt,dt_test)
	Protocol string `json:"protocol"`

	// 参数解释：监听器的监听端口。  约束限制： - QUIC监听器端口不能是4789，且不能和UDP监听器端口重复。 - 传0时表示开启监听端口范围的能力，此时port_ranges为必填字段。  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt,dt,dt_test)
	ProtocolPort *int32 `json:"protocol_port,omitempty"`

	// 参数解释：监听器使用的SNI证书（带域名的服务器证书）ID列表。  约束限制： - 列表对应的所有SNI证书的域名不允许存在重复。 - 列表对应的所有SNI证书的域名总数不超过50。 - QUIC监听器仅支持RSA证书。  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt,dt,dt_test)
	SniContainerRefs *[]string `json:"sni_container_refs,omitempty"`

	// 参数解释：监听器使用的SNI证书泛域名匹配方式。  取值范围： - longest_suffix 表示最长尾缀匹配。 - wildcard 表示标准域名分级匹配。  默认取值：wildcard。
	SniMatchAlgo *string `json:"sni_match_algo,omitempty"`

	// 参数解释：标签列表
	Tags *[]Tag `json:"tags,omitempty"`

	// 参数解释：监听器使用的安全策略。  [约束限制： - 仅对HTTPS协议类型的监听器且关联LB为独享型时有效。 - QUIC监听器不支持该字段。 - 若同时设置了security_policy_id和tls_ciphers_policy，则仅security_policy_id生效。 - 加密套件的优先顺序为ecc套件、rsa套件、tls1.3协议的套件（即支持ecc又支持rsa） ](tag:hws,hws_hk,hws_eu,g42,hk_g42,ocb,tlf,ctc,hcso,sbc,tm,cmcc,dt)  [约束限制： - 仅对HTTPS协议类型的监听器有效](tag:hcso_dt)  [取值范围：tls-1-0-inherit、tls-1-0、tls-1-1、 tls-1-2、tls-1-2-strict、tls-1-2-fs、tls-1-0-with-1-3、 tls-1-2-fs-with-1-3、 hybrid-policy-1-0、tls-1-2-strict-no-cbc，默认：tls-1-0。 ](tag:hws,hws_hk,ocb,tlf,ctc,hcso,sbc,tm,cmcc,dt)  [取值范围：tls-1-0、tls-1-1、tls-1-2、 tls-1-2-strict，默认：tls-1-0。](tag:hws_eu,g42,hk_g42,hcso_dt)  [不支持tls1.3协议的套件。](tag:tm,hws_eu,g42,hk_g42) [不支持QUIC。](tag:tm,dt,dt_test)
	TlsCiphersPolicy *string `json:"tls_ciphers_policy,omitempty"`

	// 参数解释：自定义安全策略的ID。  [约束限制： - 仅对HTTPS协议类型的监听器且关联LB为独享型时有效。 - QUIC监听器不支持该字段。 - 若同时设置了security_policy_id和tls_ciphers_policy，则仅security_policy_id生效。 - 加密套件的优先顺序为ecc套件、rsa套件、tls1.3协议的套件 （即支持ecc又支持rsa） ](tag:hws,hws_hk,hws_eu,g42,hk_g42,ocb,tlf,ctc,hcso,sbc,tm,cmcc,dt)  [约束限制： - 仅对HTTPS协议类型的监听器有效](tag:hcso_dt)  [不支持tls1.3协议的套件。](tag:tm,hws_eu,g42,hk_g42)  [不支持QUIC。](tag:tm,dt,dt_test)
	SecurityPolicyId *string `json:"security_policy_id,omitempty"`

	// 参数解释：是否开启后端服务器的重试。  [约束限制： - 若关联是共享型LB，仅在protocol为HTTP、TERMINATED_HTTPS时才能传入该字段。 - 若关联是独享型LB，仅在protocol为HTTP、HTTPS和QUIC时才能传入该字段。 ](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs,dt)  [约束限制： - 仅在protocol为HTTP、HTTPS时才能传入该字段。](tag:hws_eu,hcso_dt)  取值范围：true 开启重试，false 不开启重试。默认：true。  [不支持QUIC。](tag:tm,dt,dt_test)
	EnableMemberRetry *bool `json:"enable_member_retry,omitempty"`

	// 参数解释：客户端连接空闲超时时间。在超过keepalive_timeout时长一直没有请求， 负载均衡会暂时中断当前连接，直到下一次请求时重新建立新的连接。  约束限制：共享型实例的UDP监听器不支持此字段。  取值范围： - TCP监听器[和IP监听器](tag:hws_eu)：10-4000s，默认值为300s。 - 若为HTTP/HTTPS/TERMINATED_HTTPS监听器，取值范围为（0-4000s）默认值为60s。
	KeepaliveTimeout *int32 `json:"keepalive_timeout,omitempty"`

	// 参数解释：等待客户端请求超时时间，包括两种情况： - 读取整个客户端请求头的超时时长：如果客户端未在超时时长内发送完整个请求头，则请求将被中断 - 两个连续body体的数据包到达LB的时间间隔，超出client_timeout将会断开连接。  约束限制：仅协议为HTTP/HTTPS的监听器支持该字段。  取值范围： 1-3600s。  默认取值：60s。
	ClientTimeout *int32 `json:"client_timeout,omitempty"`

	// 参数解释：等待后端服务器响应超时时间。请求转发后端服务器后，在等待超时member_timeout时长没有响应，负载均衡将终止等待，并返回 HTTP504错误码。  约束限制：仅支持协议为HTTP/HTTPS的监听器。  取值范围：1-3600s。  默认取值：60s。
	MemberTimeout *int32 `json:"member_timeout,omitempty"`

	Ipgroup *CreateListenerIpGroupOption `json:"ipgroup,omitempty"`

	// 参数解释：是否透传客户端IP地址。开启后客户端IP地址将透传到后端服务器。  [约束限制： - 仅作用于共享型LB的TCP/UDP监听器。 - 开启特性后，ELB和后端服务器之间直接使用真实的IP访问，需要确保已正确设置服务器的安全组以及访问控制策略。 - 开启特性后，不支持同一台服务器既作为后端服务器又作为客户端的场景。 - 开启特性后，不支持变更后端服务器规格。  取值范围： - 共享型LB的TCP/UDP监听器可设置为true或false，不传默认为false。 - 共享型LB的HTTP/HTTPS监听器只支持设置为true，不传默认为true。 - 独享型负载均衡器所有协议的监听器只支持设置为true，不传默认为true。 ](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs,dt)  [约束限制：只设支持置为true，不传默认为true。](tag:hws_eu,hcso_dt)
	TransparentClientIpEnable *bool `json:"transparent_client_ip_enable,omitempty"`

	// 参数解释：是否开启proxy_protocol。  约束限制：仅TLS监听器可指定，其他协议的监听器该字段不生效，proxy_protocol不开启。
	ProxyProtocolEnable *bool `json:"proxy_protocol_enable,omitempty"`

	// 参数解释：是否开启高级转发策略功能。开启高级转发策略后，支持更灵活的转发策略和转发规则设置。  开启后支持如下场景： - 转发策略的action字段支持指定为REDIRECT_TO_URL, FIXED_RESPONSE，即支持URL重定向和响应固定的内容给客户端。 - 转发策略支持指定priority、redirect_url_config、fixed_response_config字段。 - 转发规则rule的type可以指定METHOD, HEADER, QUERY_STRING, SOURCE_IP这几种取值。 - 转发规则rule的type为HOST_NAME时，转发规则rule的value支持通配符*。 - 转发规则支持指定conditions字段。  取值范围：true开启，false不开启，默认false。  [荷兰region不支持该字段，请勿使用。](tag:dt,dt_test) > 开启后不支持关闭
	EnhanceL7policyEnable *bool `json:"enhance_l7policy_enable,omitempty"`

	QuicConfig *CreateListenerQuicConfigOption `json:"quic_config,omitempty"`

	// 参数解释：修改保护状态。  取值范围： - nonProtection: 不保护，默认值为nonProtection - consoleProtection: 控制台修改保护
	ProtectionStatus *CreateListenerOptionProtectionStatus `json:"protection_status,omitempty"`

	// 参数解释：设置保护的原因  约束限制：仅当protection_status为consoleProtection时有效。
	ProtectionReason *string `json:"protection_reason,omitempty"`

	// 参数解释：ELB是否开启gzip压缩。  [约束限制：仅HTTP/HTTPS类型监听器支持配置。](tag:tm,hws_eu,g42,hk_g42,hcso_dt,dt,dt_test) [约束限制：仅HTTP/HTTPS/QUIC类型监听器支持配置。](tag:hws,hws_hk,hws_test,hcs,hcs_sm,hcso,hk_vdf,fcs,fcs_vm,mix,ocb,ctc,cmcc,sbc,hws_ocb,hk_sbc)  取值范围：true，false  默认取值：false
	GzipEnable *bool `json:"gzip_enable,omitempty"`

	// 参数解释：端口监听范围（闭区间)，最多指定10个端口组，每个组范围不可有重叠部分。  约束限制： - 仅当protocol_port为0或未传入protoco_port时可以传入该字段。 - 仅TCP, UDP，TLS监听支持该字段
	PortRanges *[]PortRange `json:"port_ranges,omitempty"`

	// 参数解释：监听器0-RTT能力开关。  约束限制：仅HTTPS类型监听器支持配置，需要依赖TLSv1.3安全策略协议。开启 ssl_early_data 后，有重放攻击的风险，请谨慎开启。  取值范围：true，false  默认取值：false
	SslEarlyDataEnable *bool `json:"ssl_early_data_enable,omitempty"`
}

func (o CreateListenerOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerOption struct{}"
	}

	return strings.Join([]string{"CreateListenerOption", string(data)}, " ")
}

type CreateListenerOptionProtectionStatus struct {
	value string
}

type CreateListenerOptionProtectionStatusEnum struct {
	NON_PROTECTION     CreateListenerOptionProtectionStatus
	CONSOLE_PROTECTION CreateListenerOptionProtectionStatus
}

func GetCreateListenerOptionProtectionStatusEnum() CreateListenerOptionProtectionStatusEnum {
	return CreateListenerOptionProtectionStatusEnum{
		NON_PROTECTION: CreateListenerOptionProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: CreateListenerOptionProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c CreateListenerOptionProtectionStatus) Value() string {
	return c.value
}

func (c CreateListenerOptionProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateListenerOptionProtectionStatus) UnmarshalJSON(b []byte) error {
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
