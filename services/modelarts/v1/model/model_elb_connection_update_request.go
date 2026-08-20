package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ElbConnectionUpdateRequest struct {

	// **参数解释：** 负载均衡器ID。 **约束限制：**  不涉及。 **取值范围：** 不涉及。 **默认取值：**  不涉及。
	ElbId string `json:"elb_id"`

	// **参数解释：** 负载均衡器的HTTPS监听器是否开启双向认证。 **约束限制：** 仅推理服务协议为HTTPS或WSS时可配置为true，否则忽略该配置 **取值范围：** 不涉及。 **默认取值：** false
	MTls *bool `json:"m_tls,omitempty"`

	// **参数解释：** 负载均衡器的HTTPS监听器配置的客户端证书ID。 **约束限制：** 仅推理服务协议为HTTPS或WSS时可配置，否则忽略该配置 **取值范围：** 不涉及。 **默认取值：**  不涉及。
	CaCertId *string `json:"ca_cert_id,omitempty"`

	// **参数解释：** 负载均衡器的HTTPS监听器配置的服务端证书ID。 **约束限制：** 仅推理服务协议为HTTPS或WSS时可配置，否则忽略该配置。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ServerCertId *string `json:"server_cert_id,omitempty"`

	// **参数解释：** 负载均衡器的HTTPS监听器配置的SNI（服务器名称指示）证书ID列表。 **约束限制：** 仅推理服务协议为HTTPS或WSS时可配置，否则忽略该配置。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	SniCertIds *[]string `json:"sni_cert_ids,omitempty"`

	// **参数解释：** ELB接入操作。 **约束限制：** 不涉及。 **取值范围：** - RETRY：重试连接（仅当连接状态为CONNECT_FAILED时可用）。 - SYNC_ADDR：同步ELB地址（仅当连接状态为CONNECTED时可用）。 **默认取值：** 不涉及。
	Action *string `json:"action,omitempty"`
}

func (o ElbConnectionUpdateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ElbConnectionUpdateRequest struct{}"
	}

	return strings.Join([]string{"ElbConnectionUpdateRequest", string(data)}, " ")
}
