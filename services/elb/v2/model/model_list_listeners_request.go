package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListListenersRequest struct {

	// 分页查询中每页的监听器个数
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 分页查询的起始的资源id，表示上一页最后一条查询记录的负载均衡器的id。不指定时表示查询第一页。
	Marker *string `json:"marker,omitempty" xml:"marker"`

	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。
	PageReverse *bool `json:"page_reverse,omitempty" xml:"page_reverse"`

	// 监听器ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 监听器名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 监听器的描述信息。
	Description *string `json:"description,omitempty" xml:"description"`

	// 监听器所在的负载均衡器ID。
	LoadbalancerId *string `json:"loadbalancer_id,omitempty" xml:"loadbalancer_id"`

	// 监听器的最大连接数。
	ConnectionLimit *int32 `json:"connection_limit,omitempty" xml:"connection_limit"`

	// 监听器的管理状态。该字段为预留字段，暂未启用。默认为true。
	AdminStateUp *bool `json:"admin_state_up,omitempty" xml:"admin_state_up"`

	// 监听器的默认后端云服务器组ID。
	DefaultPoolId *string `json:"default_pool_id,omitempty" xml:"default_pool_id"`

	// 监听器使用的服务器证书ID。
	DefaultTlsContainerRef *string `json:"default_tls_container_ref,omitempty" xml:"default_tls_container_ref"`

	// 监听器使用的CA证书ID。
	ClientCaTlsContainerRef *string `json:"client_ca_tls_container_ref,omitempty" xml:"client_ca_tls_container_ref"`

	// 监听器的监听协议。取值范围：TCP、HTTP、UDP、TERMINATED_HTTPS。
	Protocol *string `json:"protocol,omitempty" xml:"protocol"`

	// 监听器的监听端口。
	ProtocolPort *int32 `json:"protocol_port,omitempty" xml:"protocol_port"`

	// 监听器使用的安全策略，仅对TERMINATED_HTTPS协议类型的监听器有效，且默认值为tls-1-0。取值包括：tls-1-0, tls-1-1, tls-1-2, tls-1-2-strict四种安全策略。
	TlsCiphersPolicy *string `json:"tls_ciphers_policy,omitempty" xml:"tls_ciphers_policy"`

	// 查询证书所关联的监听器
	TlsContainerId *string `json:"tls_container_id,omitempty" xml:"tls_container_id"`

	// HTTP2功能的开启状态。取值范围：true/false。true：开启。false：关闭。
	Http2Enable *bool `json:"http2_enable,omitempty" xml:"http2_enable"`

	// 企业项目ID，仅用于基于企业项目的细粒度鉴权使用。 - 如果参数传递default_pool_id，则以pool对应的企业项目ID鉴权。 - 如果default_pool_id和enterprise_project_id都没有传递 ，则进行细粒度鉴权 ，必须在用户
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`
}

func (o ListListenersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListListenersRequest struct{}"
	}

	return strings.Join([]string{"ListListenersRequest", string(data)}, " ")
}
