package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新监听器的结构体
type UpdateListenerReq struct {
	// 监听器名称。

	Name *string `json:"name,omitempty"`
	// 监听器的描述信息

	Description *string `json:"description,omitempty"`
	// 监听器的最大连接数。该字段为预留字段，暂未启用。默认为-1。

	ConnectionLimit *int32 `json:"connection_limit,omitempty"`
	// HTTP2功能的开启状态。该字段只有当监听器的协议是TERMINATED_HTTPS时才有意义。

	Http2Enable *bool `json:"http2_enable,omitempty"`
	// 监听器的默认后端云服务器组ID。当请求没有匹配的转发策略时，转发到默认后端云服务器上处理。当该字段为null时，表示监听器无默认的后端云服务器组。

	DefaultPoolId *string `json:"default_pool_id,omitempty"`
	// 监听器使用的服务器证书ID。当protocol参数为TERMINATED_HTTPS时，为必选字段

	DefaultTlsContainerRef *string `json:"default_tls_container_ref,omitempty"`
	// 监听器使用的CA证书ID。

	ClientCaTlsContainerRef *string `json:"client_ca_tls_container_ref,omitempty"`
	// 监听器使用的SNI证书（带域名的服务器证书）ID的列表。

	SniContainerRefs *[]string `json:"sni_container_refs,omitempty"`

	InsertHeaders *InsertHeader `json:"insert_headers,omitempty"`
	// 监听器使用的安全策略，仅对TERMINATED_HTTPS协议类型的监听器有效。  取值包括：tls-1-0, tls-1-1, tls-1-2, tls-1-2-strict多种安全策略。

	TlsCiphersPolicy *string `json:"tls_ciphers_policy,omitempty"`
	// 监听器的管理状态。  该字段为预留字段，暂未启动。只支持设定为true

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
}

func (o UpdateListenerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateListenerReq struct{}"
	}

	return strings.Join([]string{"UpdateListenerReq", string(data)}, " ")
}
