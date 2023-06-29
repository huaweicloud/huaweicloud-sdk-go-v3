package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListListenersRequest Request Object
type ListListenersRequest struct {

	// 分页查询中每页的监听器个数
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询的起始的资源id，表示上一页最后一条查询记录的监听器的id。不指定时表示查询第一页。
	Marker *string `json:"marker,omitempty"`

	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。
	PageReverse *bool `json:"page_reverse,omitempty"`

	// 监听器ID。
	Id *string `json:"id,omitempty"`

	// 监听器名称。
	Name *string `json:"name,omitempty"`

	// 监听器的描述信息。
	Description *string `json:"description,omitempty"`

	// 监听器所在的负载均衡器ID。
	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`

	// 监听器的最大连接数。
	ConnectionLimit *int32 `json:"connection_limit,omitempty"`

	// 监听器的管理状态。该字段为预留字段，暂未启用。默认为true。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

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

	// 查询证书所关联的监听器
	TlsContainerId *string `json:"tls_container_id,omitempty"`

	// HTTP2功能的开启状态。取值范围：true/false。true：开启。false：关闭。
	Http2Enable *bool `json:"http2_enable,omitempty"`

	// 企业项目ID。  传入all_granted_eps表示查询所有有权限的企业项目资源；\"0\"表示查询默认企业项目资源；或者指定的企业项目ID下的资源。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListListenersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListListenersRequest struct{}"
	}

	return strings.Join([]string{"ListListenersRequest", string(data)}, " ")
}
