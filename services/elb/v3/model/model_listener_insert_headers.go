package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListenerInsertHeaders struct {

	// **参数解释**：X-Forwarded-ELB-IP设为true可以将ELB实例的eip地址从报文的http头中带到后端服务器。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	XForwardedELBIP *bool `json:"X-Forwarded-ELB-IP,omitempty"`

	// **参数解释**：X-Forwarded-Port设为true可以将ELB实例的监听端口从报文的http头中带到后端服务器。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	XForwardedPort *bool `json:"X-Forwarded-Port,omitempty"`

	// **参数解释**：X-Forwarded-For-Port设为true可以将客户端的源端口从报文的http头中带到后端服务器。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	XForwardedForPort *bool `json:"X-Forwarded-For-Port,omitempty"`

	// **参数解释**：X-Forwarded-Host设为true可以将客户请求头的X-Forwarded-Host设置为请求头的Host带到后端服务器。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	XForwardedHost *bool `json:"X-Forwarded-Host,omitempty"`

	// **参数解释**：X-Forwarded-Proto设为true可以将负载均衡器实例的监听协议通过报文的http头带到后端服务器。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	XForwardedProto *bool `json:"X-Forwarded-Proto,omitempty"`

	// **参数解释**：X-Real-IP设为true可以将客户端的IP通过报文的http头带到后端服务器。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	XRealIP *bool `json:"X-Real-IP,omitempty"`

	// **参数解释**：X-Forwarded-ELB-ID设为true可以将负载均衡器实例的ID通过报文的http头带到后端服务器。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	XForwardedELBID *bool `json:"X-Forwarded-ELB-ID,omitempty"`

	// **参数解释**：X-Forwarded-TLS-Certificate-ID设为true可以将负载均衡器实例的证书ID通过报文的http头带到后端服务器。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	XForwardedTLSCertificateID *bool `json:"X-Forwarded-TLS-Certificate-ID,omitempty"`

	// **参数解释**：X-Forwarded-TLS-Protocol设为true可以将负载均衡器实例的算法协议通过报文的http头带到后端服务器。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	XForwardedTLSProtocol *bool `json:"X-Forwarded-TLS-Protocol,omitempty"`

	// **参数解释**：X-Forwarded-TLS-Cipher设为true可以将负载均衡器实例的算法套件通过报文的http头带到后端服务器。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	XForwardedTLSCipher *bool `json:"X-Forwarded-TLS-Cipher,omitempty"`
}

func (o ListenerInsertHeaders) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListenerInsertHeaders struct{}"
	}

	return strings.Join([]string{"ListenerInsertHeaders", string(data)}, " ")
}
