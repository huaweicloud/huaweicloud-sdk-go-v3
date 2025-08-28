package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMasterSlaveHealthMonitorOption **参数解释**：创建健康检查请求参数。  **约束限制**：不涉及
type CreateMasterSlaveHealthMonitorOption struct {

	// **参数解释**：健康检查间隔。  **约束限制**：不涉及  **取值范围**：1-50，单位：秒。  **默认取值**：不涉及
	Delay int32 `json:"delay"`

	// **参数解释**：发送健康检查请求的域名。  **约束限制**：当type为HTTP/HTTPS时生效。  **取值范围**：以数字或字母开头，只能包含数字、字母、’-’、’.’。  **默认取值**：null，表示使用负载均衡器的vip作为http请求的目的地址。
	DomainName *string `json:"domain_name,omitempty"`

	// **参数解释**：期望响应状态码。  **约束限制**：该字段仅在HTTP/HTTPS/GRPC协议下有效，其他协议可以设置但不会生效。  **取值范围**： - 单值：单个返回码，例如200。 - 列表：多个特定返回码，例如200，202。 - 区间：一个返回码区间，例如200-204  **默认取值**：200
	ExpectedCodes *string `json:"expected_codes,omitempty"`

	// **参数解释**：HTTP请求方法。  **约束限制**：当type为HTTP/HTTPS时生效。  **取值范围**：GET、HEAD、POST  **默认取值**：GET
	HttpMethod *string `json:"http_method,omitempty"`

	// **参数解释**：健康检查连续成功多少次后，将后端服务器的健康检查状态由OFFLINE判定为ONLINE。  **约束限制**：不涉及  **取值范围**：1-10  **默认取值**：不涉及
	MaxRetries int32 `json:"max_retries"`

	// **参数解释**：健康检查连续失败多少次后，将后端服务器的健康检查状态由ONLINE判定为OFFLINE。  **约束限制**：不涉及  **取值范围**：1-10  **默认取值**：3
	MaxRetriesDown *int32 `json:"max_retries_down,omitempty"`

	// **参数解释**：健康检查端口号。  **约束限制**：不涉及  **取值范围**：1-65535  **默认取值**：null，表示使用后端服务器端口号。
	MonitorPort *int32 `json:"monitor_port,omitempty"`

	// **参数解释**：健康检查名称。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**：一次健康检查请求的超时时间。建议该值小于delay的值。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Timeout int32 `json:"timeout"`

	// **参数解释**：健康检查请求协议。  **约束限制**： - 若pool的protocol为QUIC，则type只能是UDP_CONNECT。 - 若pool的protocol为UDP，则type只能UDP_CONNECT。 - 若pool的protocol为TCP，则type可以是TCP、HTTP、HTTPS。 - 若pool的protocol为HTTP，则type可以是TCP、HTTP、HTTPS。 - 若pool的protocol为HTTPS，则type可以是TCP、HTTP、HTTPS。 [- 不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt) [- 荷兰region不支持QUIC。](tag:dt)  **取值范围**：TCP、UDP_CONNECT、HTTP、HTTPS。  **默认取值**：不涉及
	Type string `json:"type"`

	// **参数解释**：健康检查请求的请求路径，以\"/\"开头。  **约束限制**：当type为HTTP/HTTPS时生效。  **取值范围**：支持使用字母、数字和短划线（-）、正斜线（/）、半角句号（.）、百分号（%）、半角问号（?）、井号（#）和and（&）以及扩展字符集_;~!()*[]@$^:',+  **默认取值**：\"/\"
	UrlPath *string `json:"url_path,omitempty"`
}

func (o CreateMasterSlaveHealthMonitorOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMasterSlaveHealthMonitorOption struct{}"
	}

	return strings.Join([]string{"CreateMasterSlaveHealthMonitorOption", string(data)}, " ")
}
