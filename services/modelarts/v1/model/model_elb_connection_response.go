package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ElbConnectionResponse struct {

	// **参数解释：** 负载均衡器ID。 **约束限制：**  不涉及。 **取值范围：** 不涉及。 **默认取值：**  不涉及。
	ElbId string `json:"elb_id"`

	// **参数解释：** 负载均衡器的监听器ID。 **约束限制：**  不涉及。 **取值范围：** 不涉及。 **默认取值：**  不涉及。
	ListenerId *string `json:"listener_id,omitempty"`

	// **参数解释：** 负载均衡器的HTTPS监听器是否开启双向认证。 **约束限制：** 仅推理服务协议为HTTPS或WSS时可配置为true，否则忽略该配置 **取值范围：** 不涉及。 **默认取值：** false
	MTls *bool `json:"m_tls,omitempty"`

	// **参数解释：** 负载均衡器的HTTPS监听器配置的客户端证书ID。 **约束限制：** 仅推理服务协议为HTTPS或WSS时可配置，否则忽略该配置 **取值范围：** 不涉及。 **默认取值：**  不涉及。
	CaCertId *string `json:"ca_cert_id,omitempty"`

	// **参数解释：** 负载均衡器的HTTPS监听器配置的服务端证书ID。 **约束限制：** 仅推理服务协议为HTTPS或WSS时可配置，否则忽略该配置。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ServerCertId *string `json:"server_cert_id,omitempty"`

	// **参数解释：** 负载均衡器的HTTPS监听器配置的SNI（服务器名称指示）证书ID列表。 **约束限制：** 仅推理服务协议为HTTPS或WSS时可配置，否则忽略该配置。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	SniCertIds *[]string `json:"sni_cert_ids,omitempty"`

	// **参数解释：** ELB连接状态。 **约束限制：** 仅在响应中返回，请求中传入不生效。 **取值范围：** - CONNECTING：连接中。 - CONNECTED：已连接。 - CONNECT_FAILED：连接失败。 - DISCONNECTING：断开中 - DISCONNET_FAILED：断开失败 **默认取值：** 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释：** ELB连接失败时的错误信息。 **约束限制：** 仅在响应中返回，请求中传入不生效；仅当ELB连接状态为CONNECT_FAILED时返回。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Message *string `json:"message,omitempty"`

	ResidualResources *ResidualResources `json:"residual_resources,omitempty"`
}

func (o ElbConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ElbConnectionResponse struct{}"
	}

	return strings.Join([]string{"ElbConnectionResponse", string(data)}, " ")
}
