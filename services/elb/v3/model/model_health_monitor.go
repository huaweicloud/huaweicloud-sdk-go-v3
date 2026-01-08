package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HealthMonitor struct {

	// 健康检查的管理状态。  **取值范围**： - true：表示开启健康检查。 - false表示关闭健康检查。
	AdminStateUp bool `json:"admin_state_up"`

	// **参数解释**：健康检查间隔。  **取值范围**：1-50，单位：秒。
	Delay int32 `json:"delay"`

	// **参数解释**：发送健康检查请求的域名。  **取值范围**：以数字或字母开头，只能包含数字、字母、’-’、’.’。
	DomainName string `json:"domain_name"`

	// **参数解释**：期望响应状态码。  **取值范围**： - 单值：单个返回码，例如200。 - 列表：多个特定返回码，例如200，202。 - 区间：一个返回码区间，例如200-204。
	ExpectedCodes string `json:"expected_codes"`

	// **参数解释**：HTTP请求方法。  **取值范围**：GET、HEAD、POST
	HttpMethod string `json:"http_method"`

	// **参数解释**：健康检查ID。  **取值范围**：不涉及
	Id string `json:"id"`

	// **参数解释**：健康检查连续成功多少次后，将后端服务器的健康检查状态由OFFLINE判定为ONLINE。  **取值范围**：1-10
	MaxRetries int32 `json:"max_retries"`

	// **参数解释**：健康检查连续失败多少次后，将后端服务器的健康检查状态由ONLINE判定为OFFLINE。  **取值范围**：1-10
	MaxRetriesDown int32 `json:"max_retries_down"`

	// **参数解释**：健康检查端口号。  **取值范围**：1-65535和null，传null表示使用后端服务器端口号。
	MonitorPort int32 `json:"monitor_port"`

	// **参数解释**：健康检查名称。  **取值范围**：不涉及
	Name string `json:"name"`

	// **参数解释**：健康检查所在的后端服务器组ID列表。实际只会有一个后端服务器组ID。  **取值范围**：不涉及
	Pools []PoolRef `json:"pools"`

	// **参数解释**：健康检查所在的项目ID。  **取值范围**：不涉及
	ProjectId string `json:"project_id"`

	// **参数解释**：一次健康检查请求的超时时间。  **取值范围**：不涉及
	Timeout int32 `json:"timeout"`

	// **参数解释**： 健康检查请求协议。  **取值范围**：TCP、UDP_CONNECT、HTTP、HTTPS、TLS、GRPC[和GRPCS](tag:not_open)。
	Type string `json:"type"`

	// **参数解释**：健康检查请求的请求路径。以\"/\"开头，默认为\"/\"。  **取值范围**：支持使用字母、数字和短划线（-）、正斜线（/）、半角句号（.）、百分号（%）、半角问号（?）、井号（#）和and（&）以及扩展字符集_;~!()*[]@$^:',+
	UrlPath string `json:"url_path"`

	// **参数解释**：创建时间。  **取值范围**：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。  [注意：独享型实例的历史数据以及共享型实例下的资源，不返回该字段。 ](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,srg,fcs,dt,hk_tm)
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释**： 更新时间。  **取值范围**：格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。  [注意：独享型实例的历史数据以及共享型实例下的资源，不返回该字段。 ](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,srg,fcs,dt,hk_tm)
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o HealthMonitor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthMonitor struct{}"
	}

	return strings.Join([]string{"HealthMonitor", string(data)}, " ")
}
