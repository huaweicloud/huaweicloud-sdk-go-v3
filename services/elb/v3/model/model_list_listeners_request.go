package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListListenersRequest struct {
	// 每页返回的个数。

	Limit *int32 `json:"limit,omitempty"`
	// 上一页最后一条记录的ID。  使用说明： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。

	Marker *string `json:"marker,omitempty"`
	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。  使用说明： - 必须与limit一起使用。

	PageReverse *bool `json:"page_reverse,omitempty"`
	// 监听器的前端监听端口。  支持多值查询，查询条件格式：*protocol_port=xxx&protocol_port=xxx*。

	ProtocolPort *[]int32 `json:"protocol_port,omitempty"`
	// 监听器的监听协议。  [取值：TCP、UDP、HTTP、HTTPS、TERMINATED_HTTPS。  说明：TERMINATED_HTTPS为共享型LB上的监听器独有的协议。](tag:hws,hws_hk,ocb,tlf,ctc,hcso,sbc,g42,tm,cmcc,hk-g42) [取值：TCP、UDP、HTTP、HTTPS。](tag:dt,dt_test,hcso_dt)  支持多值查询，查询条件格式：*protocol=xxx&protocol=xxx*。

	Protocol *[]string `json:"protocol,omitempty"`
	// 监听器的描述信息。  支持多值查询，查询条件格式：*description=xxx&description=xxx*。

	Description *[]string `json:"description,omitempty"`
	// 监听器的服务器证书ID。 支持多值查询，查询条件格式：*default_tls_container_ref=xxx&default_tls_container_ref=xxx*。

	DefaultTlsContainerRef *[]string `json:"default_tls_container_ref,omitempty"`
	// 监听器的CA证书ID。 支持多值查询，查询条件格式：*client_ca_tls_container_ref=xxx&client_ca_tls_container_ref=xxx*。

	ClientCaTlsContainerRef *[]string `json:"client_ca_tls_container_ref,omitempty"`
	// 监听器的管理状态，只能设置为true。  不支持该字段，请勿使用。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// ​监听器的最大连接数。取值：-1表示不限制连接数。  支持多值查询，查询条件格式：*connection_limit=xxx&connection_limit=xxx*。  不支持该字段，请勿使用。

	ConnectionLimit *[]int32 `json:"connection_limit,omitempty"`
	// 监听器的默认后端云服务器组ID。当请求没有匹配的转发策略时，转发到默认后端云服务器上处理。  支持多值查询，查询条件格式：*default_pool_id=xxx&default_pool_id=xxx*。

	DefaultPoolId *[]string `json:"default_pool_id,omitempty"`
	// 监听器ID。  支持多值查询，查询条件格式：*id=xxx&id=xxx*。

	Id *[]string `json:"id,omitempty"`
	// 监听器名称。  支持多值查询，查询条件格式：*name=xxx&name=xxx*。

	Name *[]string `json:"name,omitempty"`
	// 客户端与监听器之间的HTTPS请求的HTTP2功能的开启状态。开启后，可提升客户端与LB间的访问性能，但LB与后端服务器间仍采用HTTP1.X协议。 非HTTPS协议的监听器该字段无效，无论取值如何都不影响监听器正常运行。

	Http2Enable *bool `json:"http2_enable,omitempty"`
	// 监听器所属的负载均衡器ID。  支持多值查询，查询条件格式：*loadbalancer_id=xxx&loadbalancer_id=xxx*。

	LoadbalancerId *[]string `json:"loadbalancer_id,omitempty"`
	// 监听器使用的安全策略，仅对HTTPS协议类型的监听器有效。  [取值：tls-1-0-inherit, tls-1-0, tls-1-1, tls-1-2, tls-1-2-strict，tls-1-2-fs, tls-1-0-with-1-3, tls-1-2-fs-with-1-3, hybrid-policy-1-0。](tag:hws,hws_hk,ocb,tlf,ctc,hcso,sbc,g42,tm,cmcc,hk-g42)  [取值：tls-1-0, tls-1-1, tls-1-2, tls-1-2-strict。](tag:dt,dt_test)  支持多值查询，查询条件格式：*tls_ciphers_policy=xxx&tls_ciphers_policy=xxx*。

	TlsCiphersPolicy *[]string `json:"tls_ciphers_policy,omitempty"`
	// 后端云服务器的IP地址。仅用于查询条件，不作为响应参数字段。  支持多值查询，查询条件格式：*member_address=xxx&member_address=xxx*。

	MemberAddress *[]string `json:"member_address,omitempty"`
	// 后端云服务器对应的弹性云服务器的ID。仅用于查询条件，不作为响应参数字段。  支持多值查询，查询条件格式：*member_device_id=xxx&member_device_id=xxx*。

	MemberDeviceId *[]string `json:"member_device_id,omitempty"`
	// 企业项目ID。  支持多值查询，查询条件格式：*enterprise_project_id=xxx&enterprise_project_id=xxx*。  [不支持该字段，请勿使用。](tag:dt,dt_test,hcso_dt)

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
	// 是否开启后端服务器的重试。取值：true 开启重试，false 不开启重试。

	EnableMemberRetry *bool `json:"enable_member_retry,omitempty"`
	// 等待后端服务器响应超时时间。请求转发后端服务器后，在等待超时member_timeout时长没有响应，负载均衡将终止等待，并返回 HTTP504错误码。  取值：1-300s。  支持多值查询，查询条件格式：*member_timeout=xxx&member_timeout=xxx*。

	MemberTimeout *[]int32 `json:"member_timeout,omitempty"`
	// 等待客户端请求超时时间，包括两种情况： - 读取整个客户端请求头的超时时长：如果客户端未在超时时长内发送完整个请求头，则请求将被中断 - 两个连续body体的数据包到达LB的时间间隔，超出client_timeout将会断开连接。  取值：1-300s。  支持多值查询，查询条件格式：*client_timeout=xxx&client_timeout=xxx*。

	ClientTimeout *[]int32 `json:"client_timeout,omitempty"`
	// 客户端连接空闲超时时间。在超过keepalive_timeout时长一直没有请求，负载均衡会暂时中断当前连接，直到一下次请求时重新建立新的连接。取值：  - TCP监听器：10-4000s。  - HTTP/HTTPS/TERMINATED_HTTPS监听器：0-4000s。  - UDP监听器不支持此字段。 支持多值查询，查询条件格式：*keepalive_timeout=xxx&keepalive_timeout=xxx*。

	KeepaliveTimeout *[]int32 `json:"keepalive_timeout,omitempty"`
	// 是否透传客户端IP地址。 [开启后客户端IP地址将透传到后端服务器。仅作用于共享型LB的TCP/UDP监听器。取值：true开启，false不开启。](tag:hws,hws_hk,ocb,tlf,ctc,hcso,sbc,g42,tm,cmcc,hk-g42,dt,dt_test) [不支持该字段，请勿使用。](tag:hcso_dt)

	TransparentClientIpEnable *bool `json:"transparent_client_ip_enable,omitempty"`
	// 是否开启高级转发策略功能。开启高级转发策略后，支持更灵活的转发策略和转发规则设置。取值：true开启，false不开启。

	EnhanceL7policyEnable *bool `json:"enhance_l7policy_enable,omitempty"`
	// 后端云服务器ID。仅用于查询条件，不作为响应参数字段。  支持多值查询，查询条件格式：*member_instance_id=xxx&member_instance_id=xxx*。

	MemberInstanceId *[]string `json:"member_instance_id,omitempty"`
}

func (o ListListenersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListListenersRequest struct{}"
	}

	return strings.Join([]string{"ListListenersRequest", string(data)}, " ")
}
