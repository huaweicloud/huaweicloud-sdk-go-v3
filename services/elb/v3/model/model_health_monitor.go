package model

import (
	"encoding/json"

	"strings"
)

// 健康检查对象
type HealthMonitor struct {
	// 功能说明：管理状态true/false。true表示开启健康检查，false表示关闭健康检查。

	AdminStateUp bool `json:"admin_state_up"`
	// 健康检查间隔

	Delay int32 `json:"delay"`
	// 功能说明：健康检查测试member健康状态时，发送的http请求的域名。仅当type为HTTP时生效。使用说明：默认为空，表示使用负载均衡器的vip作为http请求的目的地址。以数字或字母开头，只能包含数字、字母、’-’、’.’。

	DomainName string `json:"domain_name"`
	// 期望HTTP响应状态码，指定下列值：单值，例如200；列表，例如200，202；区间，例如200-204。仅当type为HTTP时生效。该字段为预留字段，暂未启用。

	ExpectedCodes string `json:"expected_codes"`
	// HTTP方法，可以为GET、HEAD、POST、PUT、DELETE、TRACE、OPTIONS、CONNECT、PATCH。仅当type为HTTP时生效。该字段为预留字段，暂未启用。

	HttpMethod string `json:"http_method"`
	// 健康检查ID

	Id string `json:"id"`
	// 健康检查连续成功多少次后，将后端服务器的健康检查状态由offline判定为online，取值范围[1，10]。

	MaxRetries int32 `json:"max_retries"`
	// 健康检查连续成功多少次后，将后端服务器的健康检查状态由ONLINE判定为OFFLINE

	MaxRetriesDown int32 `json:"max_retries_down"`
	// 健康检查端口号。默认为空，表示使用后端云服务器组的端口。

	MonitorPort int32 `json:"monitor_port"`
	// 健康检查名称。

	Name string `json:"name"`
	// 健康检查关联的后端云服务器组列表

	Pools []PoolRef `json:"pools"`
	// 健康检查所在的项目ID。

	ProjectId string `json:"project_id"`
	// 健康检查的超时时间。建议该值小于delay的值。

	Timeout int32 `json:"timeout"`
	// 健康检查类型

	Type string `json:"type"`
	// 功能说明：健康检查测试member健康时发送的http请求路径。使用说明：以“/”开头。仅当type为HTTP时生效。

	UrlPath string `json:"url_path"`
}

func (o HealthMonitor) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "HealthMonitor struct{}"
	}

	return strings.Join([]string{"HealthMonitor", string(data)}, " ")
}
