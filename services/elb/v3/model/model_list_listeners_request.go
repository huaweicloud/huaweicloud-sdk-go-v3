package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListListenersRequest Request Object
type ListListenersRequest struct {

	// **参数解释**：每页返回的个数。  **约束限制**：不涉及  **取值范围**：0-2000  **默认取值**：2000
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：上一页最后一条记录的ID。  **约束限制**： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。  **取值范围**：不涉及  **默认取值**：不涉及
	Marker *string `json:"marker,omitempty"`

	// **参数解释**：是否反向查询。  **约束限制**： - 必须与limit一起使用。 - 当page_reverse=true时，若要查询上一页，marker取值为当前页返回值的previous_marker。  **取值范围**： - true：查询上一页。 - false：查询下一页。  **默认取值**：false
	PageReverse *bool `json:"page_reverse,omitempty"`

	// **参数解释**：监听器的前端监听端口。 支持多值查询，查询条件格式：*protocol_port=xxx&protocol_port=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	ProtocolPort *[]string `json:"protocol_port,omitempty"`

	// **参数解释**：监听器的监听协议。 支持多值查询，查询条件格式：*protocol=xxx&protocol=xxx*。  **约束限制**：不涉及  [取值范围：TCP、UDP、HTTP、HTTPS、TERMINATED_HTTPS、QUIC、TLS。 说明：TERMINATED_HTTPS为共享型LB上的监听器独有的协议。](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,srg,fcs,dt)  [取值范围：TCP、UDP、HTTP、HTTPS。](tag:hcso_dt) [取值范围：TCP、UDP、IP、HTTP、HTTPS。IP为网关型LB上的监听器独有的协议。](tag:hws_eu)  **默认取值**：不涉及  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt,dt)
	Protocol *[]string `json:"protocol,omitempty"`

	// **参数解释**：监听器的描述信息。 支持多值查询，查询条件格式：*description=xxx&description=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Description *[]string `json:"description,omitempty"`

	// **参数解释**：监听器的服务器证书ID。 支持多值查询，查询条件格式： *default_tls_container_ref=xxx&default_tls_container_ref=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	DefaultTlsContainerRef *[]string `json:"default_tls_container_ref,omitempty"`

	// **参数解释**：监听器的CA证书ID。 支持多值查询，查询条件格式： *client_ca_tls_container_ref=xxx&client_ca_tls_container_ref=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	ClientCaTlsContainerRef *[]string `json:"client_ca_tls_container_ref,omitempty"`

	// **参数解释**：监听器的管理状态。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及  [不支持该字段，请勿使用。](tag:dt,hcso_dt)
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// **参数解释**：查询结果是否包含回收站负载均衡器的监听器  **约束限制**：不涉及  **取值范围**： - true ：包含回收站elb的监听器。 - false：不包含回收站elb的监听器。  **默认取值**：不涉及
	IncludeRecycleBin *bool `json:"include_recycle_bin,omitempty"`

	// **参数解释**：监听器的最大连接数。 支持多值查询，查询条件格式：*connection_limit=xxx&connection_limit=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	ConnectionLimit *[]int32 `json:"connection_limit,omitempty"`

	// **参数解释**：监听器的默认后端服务器组ID。当请求没有匹配的转发策略时，转发到默认后端服务器组上处理。 支持多值查询，查询条件格式：*default_pool_id=xxx&default_pool_id=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	DefaultPoolId *[]string `json:"default_pool_id,omitempty"`

	// **参数解释**：监听器ID。 支持多值查询，查询条件格式：*id=xxx&id=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Id *[]string `json:"id,omitempty"`

	// **参数解释**：监听器名称。 支持多值查询，查询条件格式：*name=xxx&name=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Name *[]string `json:"name,omitempty"`

	// **参数解释**：客户端与LB之间的HTTPS请求的HTTP2功能的开启状态。 开启后，可提升客户端与LB间的访问性能，但LB与后端服务器间仍采用HTTP1.X协议。  **约束限制**：不涉及  **取值范围**：true 开启，false 不开启。  **默认取值**：不涉及  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt,dt)
	Http2Enable *bool `json:"http2_enable,omitempty"`

	// **参数解释**：监听器所属的负载均衡器ID。 支持多值查询，查询条件格式：*loadbalancer_id=xxx&loadbalancer_id=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	LoadbalancerId *[]string `json:"loadbalancer_id,omitempty"`

	// **参数解释**：监听器使用的安全策略。 支持多值查询，查询条件格式：*tls_ciphers_policy=xxx&tls_ciphers_policy=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	TlsCiphersPolicy *[]string `json:"tls_ciphers_policy,omitempty"`

	// **参数解释**：后端服务器的IP地址。仅用于查询条件，不作为响应参数字段。 支持多值查询，查询条件格式：*member_address=xxx&member_address=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	MemberAddress *[]string `json:"member_address,omitempty"`

	// **参数解释**：后端服务器对应的弹性云服务器的ID。仅用于查询条件，不作为响应参数字段。 支持多值查询，查询条件格式：*member_device_id=xxx&member_device_id=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	MemberDeviceId *[]string `json:"member_device_id,omitempty"`

	// **参数解释**：资源所属的企业项目ID。 支持多值查询，查询条件格式： *enterprise_project_id=xxx&enterprise_project_id=xxx*。  **约束限制**： 如果enterprise_project_id不传值，默认查询所有企业项目下的资源，鉴权按照细粒度权限鉴权，必须在用户组下分配elb:listeners:list权限。 如果enterprise_project_id传值，鉴权按照企业项目权限鉴权，分为传入具体eps_id和all_granted_eps两种场景，前者查询指定eps_id的eps下的资源，后者查询的是所有有list权限的eps下的资源。  **取值范围**：不涉及  **默认取值**：不涉及  [不支持该字段，请勿使用。](tag:dt,hcso_dt)
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// **参数解释**：是否开启后端服务器的重试。  **约束限制**：不涉及  **取值范围**：true 开启，false 不开启。  **默认取值**：不涉及
	EnableMemberRetry *bool `json:"enable_member_retry,omitempty"`

	// **参数解释**：等待后端服务器响应超时时间。请求转发后端服务器后，在等待超时member_timeout时长没有响应，负载均衡将终止等待，并返回HTTP504错误码。 支持多值查询，查询条件格式：*member_timeout=xxx&member_timeout=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	MemberTimeout *[]int32 `json:"member_timeout,omitempty"`

	// **参数解释**：等待客户端请求超时时间，包括两种情况： - 读取整个客户端请求头的超时时长：如果客户端未在超时时长内发送完整个请求头，则请求将被中断 - 两个连续body体的数据包到达LB的时间间隔，超出client_timeout将会断开连接。 支持多值查询，查询条件格式：*client_timeout=xxx&client_timeout=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	ClientTimeout *[]int32 `json:"client_timeout,omitempty"`

	// **参数解释**：客户端连接空闲超时时间。在超过keepalive_timeout时长一直没有请求，负载均衡会暂时中断当前连接，直到下一次请求时重新建立新的连接。 支持多值查询，查询条件格式：*keepalive_timeout=xxx&keepalive_timeout=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	KeepaliveTimeout *[]int32 `json:"keepalive_timeout,omitempty"`

	// **参数解释**：是否开启后客户端IP地址将透传到后端服务器。 [仅作用于共享型LB的TCP/UDP监听器。](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,srg,fcs,dt,hk_tm)  **约束限制**：不涉及  **取值范围**：true开启，false不开启。  **默认取值**：不涉及
	TransparentClientIpEnable *bool `json:"transparent_client_ip_enable,omitempty"`

	// **参数解释**：是否开启proxy_protocol。  **约束限制**：不涉及  **取值范围**：true开启，false不开启。  **默认取值**：不涉及
	ProxyProtocolEnable *bool `json:"proxy_protocol_enable,omitempty"`

	// **参数解释**：是否开启高级转发策略功能。开启高级转发策略后，支持更灵活的转发策略和转发规则设置。  **约束限制**：不涉及  **取值范围**：true开启，false不开启。  **默认取值**：不涉及  [荷兰region不支持该字段，请勿使用。](tag:dt)
	EnhanceL7policyEnable *bool `json:"enhance_l7policy_enable,omitempty"`

	// **参数解释**：后端服务器ID。仅用于查询条件，不作为响应参数字段。 支持多值查询，查询条件格式：*member_instance_id=xxx&member_instance_id=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	MemberInstanceId *[]string `json:"member_instance_id,omitempty"`

	// **参数解释**：修改保护状态。 支持多值查询，查询条件格式：*protection_status=xxx&protection_status=xxx*。  **约束限制**：不涉及  **取值范围**： - nonProtection: 不保护，默认值为nonProtection - consoleProtection: 控制台修改保护  **默认取值**：不涉及
	ProtectionStatus *[]string `json:"protection_status,omitempty"`

	// **参数解释**：是否开启监听器0-RTT功能。  **约束限制**：不涉及  **取值范围**：true开启，false不开启。  **默认取值**：不涉及
	SslEarlyDataEnable *bool `json:"ssl_early_data_enable,omitempty"`

	// **参数解释**：是否开启nat64地址转换功能。  **约束限制**：不涉及  **取值范围**：true开启，false不开启。  **默认取值**：不涉及
	Nat64Enable *bool `json:"nat64_enable,omitempty"`
}

func (o ListListenersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListListenersRequest struct{}"
	}

	return strings.Join([]string{"ListListenersRequest", string(data)}, " ")
}
