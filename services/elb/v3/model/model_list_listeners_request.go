package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListListenersRequest Request Object
type ListListenersRequest struct {

	// 参数解释：每页返回的个数。  取值范围：0-2000  默认取值：2000
	Limit *int32 `json:"limit,omitempty"`

	// 上一页最后一条记录的ID。  使用说明： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。
	Marker *string `json:"marker,omitempty"`

	// 是否反向查询。  取值： - true：查询上一页。 - false：查询下一页，默认。  使用说明： - 必须与limit一起使用。 - 当page_reverse=true时，若要查询上一页，marker取值为当前页返回值的previous_marker。
	PageReverse *bool `json:"page_reverse,omitempty"`

	// 监听器的前端监听端口。  [当监听器的protocol为IP时，前端端口固定为0。](tag:hws_eu) 支持多值查询，查询条件格式：*protocol_port=xxx&protocol_port=xxx*。
	ProtocolPort *[]string `json:"protocol_port,omitempty"`

	// 监听器的监听协议。  [取值：TCP、UDP、HTTP、HTTPS、TERMINATED_HTTPS、QUIC、TLS。 说明：TERMINATED_HTTPS为共享型LB上的监听器独有的协议。](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,fcs,dt)  [取值：TCP、UDP、HTTP、HTTPS。](tag:hcso_dt) [取值：TCP、UDP、IP、HTTP、HTTPS。IP为网关型LB上的监听器独有的协议。](tag:hws_eu)  支持多值查询，查询条件格式：*protocol=xxx&protocol=xxx*。  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt,dt,dt_test)
	Protocol *[]string `json:"protocol,omitempty"`

	// 监听器的描述信息。  支持多值查询，查询条件格式：*description=xxx&description=xxx*。
	Description *[]string `json:"description,omitempty"`

	// 监听器的服务器证书ID。  支持多值查询，查询条件格式： *default_tls_container_ref=xxx&default_tls_container_ref=xxx*。
	DefaultTlsContainerRef *[]string `json:"default_tls_container_ref,omitempty"`

	// 监听器的CA证书ID。  支持多值查询，查询条件格式： *client_ca_tls_container_ref=xxx&client_ca_tls_container_ref=xxx*。
	ClientCaTlsContainerRef *[]string `json:"client_ca_tls_container_ref,omitempty"`

	// 监听器的管理状态。  [不支持该字段，请勿使用。](tag:dt,dt_test,hcso_dt)
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// ​监听器的最大连接数。  取值：-1表示不限制连接数。  支持多值查询，查询条件格式：*connection_limit=xxx&connection_limit=xxx*。  不支持该字段，请勿使用。
	ConnectionLimit *[]int32 `json:"connection_limit,omitempty"`

	// 监听器的默认后端服务器组ID。当请求没有匹配的转发策略时，转发到默认后端服务器上处理。  支持多值查询，查询条件格式：*default_pool_id=xxx&default_pool_id=xxx*。
	DefaultPoolId *[]string `json:"default_pool_id,omitempty"`

	// 监听器ID。  支持多值查询，查询条件格式：*id=xxx&id=xxx*。
	Id *[]string `json:"id,omitempty"`

	// 监听器名称。  支持多值查询，查询条件格式：*name=xxx&name=xxx*。
	Name *[]string `json:"name,omitempty"`

	// 客户端与LB之间的HTTPS请求的HTTP2功能的开启状态。 开启后，可提升客户端与LB间的访问性能，但LB与后端服务器间仍采用HTTP1.X协议。  使用说明： - 仅HTTPS协议监听器有效。 - QUIC监听器不能设置该字段，固定返回为true。 - 其他协议的监听器可设置该字段但无效，无论取值如何都不影响监听器正常运行。  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt,dt,dt_test)
	Http2Enable *bool `json:"http2_enable,omitempty"`

	// 监听器所属的负载均衡器ID。  支持多值查询，查询条件格式：*loadbalancer_id=xxx&loadbalancer_id=xxx*。
	LoadbalancerId *[]string `json:"loadbalancer_id,omitempty"`

	// 监听器使用的安全策略。  支持多值查询，查询条件格式：*tls_ciphers_policy=xxx&tls_ciphers_policy=xxx*。
	TlsCiphersPolicy *[]string `json:"tls_ciphers_policy,omitempty"`

	// 后端服务器的IP地址。仅用于查询条件，不作为响应参数字段。  支持多值查询，查询条件格式：*member_address=xxx&member_address=xxx*。
	MemberAddress *[]string `json:"member_address,omitempty"`

	// 后端服务器对应的弹性云服务器的ID。仅用于查询条件，不作为响应参数字段。  支持多值查询，查询条件格式：*member_device_id=xxx&member_device_id=xxx*。
	MemberDeviceId *[]string `json:"member_device_id,omitempty"`

	// 企业项目ID。不传时查询default企业项目\"0\"下的资源，鉴权按照default企业项目鉴权； 如果传值，则传已存在的企业项目ID或all_granted_eps（表示查询所有企业项目）进行查询。  支持多值查询，查询条件格式：*enterprise_project_id=xxx&enterprise_project_id=xxx*。  [不支持该字段，请勿使用。](tag:dt,dt_test,hcso_dt)
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// 是否开启后端服务器的重试。  取值：true 开启重试，false 不开启重试。
	EnableMemberRetry *bool `json:"enable_member_retry,omitempty"`

	// 等待后端服务器响应超时时间。请求转发后端服务器后，在等待超时member_timeout时长没有响应，负载均衡将终止等待，并返回 HTTP504错误码。  取值：1-3600s。  支持多值查询，查询条件格式：*member_timeout=xxx&member_timeout=xxx*。
	MemberTimeout *[]int32 `json:"member_timeout,omitempty"`

	// 等待客户端请求超时时间，包括两种情况： - 读取整个客户端请求头的超时时长：如果客户端未在超时时长内发送完整个请求头，则请求将被中断 - 两个连续body体的数据包到达LB的时间间隔，超出client_timeout将会断开连接。  取值：1-3600s。  支持多值查询，查询条件格式：*client_timeout=xxx&client_timeout=xxx*。
	ClientTimeout *[]int32 `json:"client_timeout,omitempty"`

	// 客户端连接空闲超时时间。在超过keepalive_timeout时长一直没有请求， 负载均衡会暂时中断当前连接，直到下一次请求时重新建立新的连接。  取值： - TCP监听器[和IP监听器](tag:hws_eu)：10-4000s。 - HTTP/HTTPS/TERMINATED_HTTPS监听器：0-4000s。 [- 共享型实例的UDP监听器不支持此字段。](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,fcs,dt,dt_test,hk_tm)  支持多值查询，查询条件格式：*keepalive_timeout=xxx&keepalive_timeout=xxx*。
	KeepaliveTimeout *[]int32 `json:"keepalive_timeout,omitempty"`

	// 是否透传客户端IP地址。开启后客户端IP地址将透传到后端服务器。  [仅作用于共享型LB的TCP/UDP监听器。取值：true开启，false不开启。 ](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs,dt,hk_tm)
	TransparentClientIpEnable *bool `json:"transparent_client_ip_enable,omitempty"`

	// 是否开启proxy_protocol。仅TLS监听器可指定，其他协议的监听器该字段不生效，proxy_protocol不开启。
	ProxyProtocolEnable *bool `json:"proxy_protocol_enable,omitempty"`

	// 是否开启高级转发策略功能。开启高级转发策略后，支持更灵活的转发策略和转发规则设置。  取值：true开启，false不开启。  [荷兰region不支持该字段，请勿使用。](tag:dt,dt_test)
	EnhanceL7policyEnable *bool `json:"enhance_l7policy_enable,omitempty"`

	// 后端服务器ID。仅用于查询条件，不作为响应参数字段。  支持多值查询，查询条件格式：*member_instance_id=xxx&member_instance_id=xxx*。
	MemberInstanceId *[]string `json:"member_instance_id,omitempty"`

	// 修改保护状态, 取值： - nonProtection: 不保护，默认值为nonProtection - consoleProtection: 控制台修改保护
	ProtectionStatus *[]string `json:"protection_status,omitempty"`
}

func (o ListListenersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListListenersRequest struct{}"
	}

	return strings.Join([]string{"ListListenersRequest", string(data)}, " ")
}
