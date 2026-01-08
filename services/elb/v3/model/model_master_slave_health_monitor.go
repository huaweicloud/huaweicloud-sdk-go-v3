package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MasterSlaveHealthMonitor 健康检查对象
type MasterSlaveHealthMonitor struct {

	// **参数解释**：健康检查的管理状态。  **取值范围**： - true：表示开启健康检查。 - false表示关闭健康检查。
	AdminStateUp bool `json:"admin_state_up"`

	// **参数解释**：健康检查间隔。  **取值范围**：1-50，单位：秒。
	Delay int32 `json:"delay"`

	// **参数解释**：发送健康检查请求的域名。  **取值范围**：以数字或字母开头，只能包含数字、字母、’-’、’.’。
	DomainName string `json:"domain_name"`

	// **参数解释**：期望响应状态码。  **取值范围**： - 单值：单个返回码，例如200。 - 列表：多个特定返回码，例如200，202。 - 区间：一个返回码区间，例如200-204。
	ExpectedCodes string `json:"expected_codes"`

	// **约束限制**：HTTP请求方法。  **取值范围**：GET、HEAD、POST
	HttpMethod string `json:"http_method"`

	// **参数解释**：健康检查ID  **取值范围**：不涉及
	Id string `json:"id"`

	// **参数解释**：健康检查连续成功多少次后，将后端服务器的健康检查状态由OFFLINE判定为ONLINE。  **取值范围**：1-10
	MaxRetries int32 `json:"max_retries"`

	// **参数解释**：健康检查连续失败多少次后，将后端服务器的健康检查状态由ONLINE判定为OFFLINE。  **取值范围**：1-10
	MaxRetriesDown int32 `json:"max_retries_down"`

	// **参数解释**：健康检查端口号。  **取值范围**：1-65535
	MonitorPort int32 `json:"monitor_port"`

	// **参数解释**：健康检查名称。  **取值范围**：不涉及
	Name string `json:"name"`

	// **参数解释**：一次健康检查请求的超时时间。  **取值范围**：不涉及
	Timeout int32 `json:"timeout"`

	// **参数解释**：健康检查请求协议。  **取值范围**：TCP、UDP_CONNECT、HTTP、HTTPS。
	Type string `json:"type"`

	// **参数解释**：健康检查请求的请求路径。以\"/\"开头，默认为\"/\"。  **取值范围**：支持使用字母、数字和短划线（-）、正斜线（/）、半角句号（.）、百分号（%）、半角问号（?）、井号（#）和and（&）以及扩展字符集_;~!()*[]@$^:',+
	UrlPath string `json:"url_path"`
}

func (o MasterSlaveHealthMonitor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MasterSlaveHealthMonitor struct{}"
	}

	return strings.Join([]string{"MasterSlaveHealthMonitor", string(data)}, " ")
}
