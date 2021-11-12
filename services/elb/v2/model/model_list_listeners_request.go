package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListListenersRequest struct {
	// 分页查询中每页的监听器个数

	Limit *int32 `json:"limit,omitempty"`
	// 分页查询的起始的资源id，表示上一页最后一条查询记录的负载均衡器的id。不指定时表示查询第一页。

	Marker *string `json:"marker,omitempty"`
	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。

	PageReverse *bool `json:"page_reverse,omitempty"`
	// 监听器ID。

	Id *string `json:"id,omitempty"`
	// 监听器名称。

	Name *string `json:"name,omitempty"`
	// 监听器的描述信息。

	Description *string `json:"description,omitempty"`
	// 监听器的默认后端云服务器组ID。

	DefaultPoolId *string `json:"default_pool_id,omitempty"`
	// 监听器使用的服务器证书ID。

	DefaultTlsContainerRef *string `json:"default_tls_container_ref,omitempty"`
	// 监听器使用的CA证书ID。

	ClientCaTlsContainerRef *string `json:"client_ca_tls_container_ref,omitempty"`
	// 监听器的监听协议。取值范围：TCP、HTTP、UDP、TERMINATED_HTTPS。

	Protocol *string `json:"protocol,omitempty"`
	// 监听器的监听端口。

	ProtocolPort *int32 `json:"protocol_port,omitempty"`
	// 监听器使用的安全策略，仅对TERMINATED_HTTPS协议类型的监听器有效，且默认值为tls-1-0。取值包括：tls-1-0, tls-1-1, tls-1-2, tls-1-2-strict四种安全策略。

	TlsCiphersPolicy *string `json:"tls_ciphers_policy,omitempty"`
	// 等待后端服务器请求超时时间，协议为HTTP， TERMINATED_HTTPS时才有意义。取值范围 1-300

	MemberTimeout *int32 `json:"member_timeout,omitempty"`
	// 等待客户端请求超时时间，协议为HTTP， TERMINATED_HTTPS的监听器才有意义。取值范围 1-60

	ClientTimeout *int32 `json:"client_timeout,omitempty"`
	// TCP监听器配置空闲超时时间，取值范围为（10-900s）默认值为300s，TCP监听器配置空闲超时时间，取值范围为（10-900s）默认值为300s，HTTP/TERMINATED_HTTPS监听器为客户端连接空闲超时时间，取值范围为（1-300s）默认值为15s。 UDP此字段无意义

	KeepaliveTimeout *int32 `json:"keepalive_timeout,omitempty"`
	// 查询证书所关联的监听器

	TlsContainerId *string `json:"tls_container_id,omitempty"`
}

func (o ListListenersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListListenersRequest struct{}"
	}

	return strings.Join([]string{"ListListenersRequest", string(data)}, " ")
}
