package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListListenersRequest struct {
	// 监听器的管理状态。只支持设定为true，该字段的值无实际意义。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 监听器使用的CA证书ID。

	ClientCaTlsContainerRef *[]string `json:"client_ca_tls_container_ref,omitempty"`
	// 等待客户端请求超时时间，协议为HTTP， TERMINATED_HTTPS的监听器才有意义。取值范围 1-60

	ClientTimeout *[]int32 `json:"client_timeout,omitempty"`
	// 监听器的最大连接数。该字段为预留字段，暂未启用。默认为-1。

	ConnectionLimit *[]int32 `json:"connection_limit,omitempty"`
	// 监听器的默认后端云服务器组ID。当请求没有匹配的转发策略时，转发到默认后端云服务器上处理。

	DefaultPoolId *[]string `json:"default_pool_id,omitempty"`
	// 监听器使用的服务器证书ID。

	DefaultTlsContainerRef *[]string `json:"default_tls_container_ref,omitempty"`
	// 监听器的描述信息。

	Description *[]string `json:"description,omitempty"`
	// 后端重试探测的开关

	EnableMemberRetry *bool `json:"enable_member_retry,omitempty"`
	// 企业项目ID。

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
	// HTTP2功能的开启状态。该字段只有当监听器的协议是TERMINATED_HTTPS时生效。

	Http2Enable *bool `json:"http2_enable,omitempty"`
	// 监听器ID。

	Id *[]string `json:"id,omitempty"`
	// TCP监听器配置空闲超时时间，取值范围为（10-900s）默认值为300s，TCP监听器配置空闲超时时间，取值范围为（10-900s）默认值为300s，HTTP/TERMINATED_HTTPS监听器为客户端连接空闲超时时间，取值范围为（1-300s）默认值为15s。 UDP此字段无意义

	KeepaliveTimeout *[]int32 `json:"keepalive_timeout,omitempty"`
	// 每页返回的个数。

	Limit *int32 `json:"limit,omitempty"`
	// 监听器绑定的负载均衡器ID。

	LoadbalancerId *[]string `json:"loadbalancer_id,omitempty"`
	// 上一页最后一条记录的ID。  使用说明：  - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。

	Marker *string `json:"marker,omitempty"`
	// 后端云服务器的IP地址。

	MemberAddress *[]string `json:"member_address,omitempty"`
	// 后端云服务器对应的弹性云服务器的ID。

	MemberDeviceId *[]string `json:"member_device_id,omitempty"`
	// 等待后端服务器请求超时时间，协议为HTTP， TERMINATED_HTTPS时才有意义。取值范围 1-300

	MemberTimeout *[]int32 `json:"member_timeout,omitempty"`
	// 监听器名称。

	Name *[]string `json:"name,omitempty"`
	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。 使用说明：必须与limit一起使用。

	PageReverse *bool `json:"page_reverse,omitempty"`
	// 监听器的监听协议。 取值：UDP,TCP,HTTP,TERMINATED_HTTPS。

	Protocol *[]string `json:"protocol,omitempty"`
	// 监听器的监听端口。

	ProtocolPort *[]string `json:"protocol_port,omitempty"`
	// 监听器使用的安全策略，仅对TERMINATED_HTTPS协议类型的监听器有效，且默认值为tls-1-0。 取值包括：tls-1-0, tls-1-1, tls-1-2, tls-1-2-strict四种安全策略。

	TlsCiphersPolicy *[]string `json:"tls_ciphers_policy,omitempty"`
	// 获取客户端真实IP

	TransparentClientIpEnable *bool `json:"transparent_client_ip_enable,omitempty"`
}

func (o ListListenersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListListenersRequest struct{}"
	}

	return strings.Join([]string{"ListListenersRequest", string(data)}, " ")
}
