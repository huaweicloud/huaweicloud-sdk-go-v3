package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateEsListenerRequestBody struct {

	// 协议类型，支持HTTP、HTTPS
	Protocol string `json:"protocol"`

	// 端口。
	ProtocolPort int32 `json:"protocol_port"`

	// server证书Id。如protocol为HTTPS则该字段必选。
	ServerCertId *string `json:"server_cert_id,omitempty"`

	// CA证书Id。如protocol为HTTPS且为双向认证时则该字段必选。
	CaCertId *string `json:"ca_cert_id,omitempty"`

	// 类型：searchTool 表示对Elasticsearch/Opensearch进行监听器配，viewTool 表示对Kibana/Opensearch Dashboard进行监听器配置，默认为searchTool 。
	Type *string `json:"type,omitempty"`
}

func (o CreateEsListenerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEsListenerRequestBody struct{}"
	}

	return strings.Join([]string{"CreateEsListenerRequestBody", string(data)}, " ")
}
