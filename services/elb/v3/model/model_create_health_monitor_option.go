package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建健康检查请求参数。
type CreateHealthMonitorOption struct {
	// 健康检查的管理状态。取值： - true：表示开启健康检查，默认为true。 - false表示关闭健康检查。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 健康检查间隔。取值：1-50s。

	Delay int32 `json:"delay"`
	// 发送健康检查请求的域名。  取值：以数字或字母开头，只能包含数字、字母、'-'、'.'。  默认为空，表示使用负载均衡器的vip作为http请求的目的地址。  使用说明： - 仅当type为HTTP时生效。

	DomainName *string `json:"domain_name,omitempty"`
	// 期望响应状态码。支持多种取值格式： - 单值：单个返回码，例如200。 - 列表：多个特定返回码，例如200，202。 - 区间：一个返回码区间，例如200-204。   默认值：200。  仅支持HTTP/HTTPS设置该字段，其他协议设置不会生效。

	ExpectedCodes *string `json:"expected_codes,omitempty"`
	// HTTP请求方法，取值：GET、HEAD、POST、PUT、DELETE、TRACE、OPTIONS、CONNECT、PATCH，默认GET。 使用说明：  - 仅当type为HTTP时生效。 不支持该字段，请勿使用。

	HttpMethod *string `json:"http_method,omitempty"`
	// 健康检查连续成功多少次后，将后端服务器的健康检查状态由OFFLINE判定为ONLINE。取值范围：1-10。

	MaxRetries int32 `json:"max_retries"`
	// 健康检查连续失败多少次后，将后端服务器的健康检查状态由ONLINE判定为OFFLINE。取值范围：1-10，默认3。

	MaxRetriesDown *int32 `json:"max_retries_down,omitempty"`
	// 健康检查端口号。取值：1-65535，默认为空，表示使用后端云服务器端口号。

	MonitorPort *int32 `json:"monitor_port,omitempty"`
	// 健康检查名称。

	Name *string `json:"name,omitempty"`
	// 健康检查所在的后端云服务器组ID

	PoolId string `json:"pool_id"`
	// 健康检查所在的项目ID。

	ProjectId *string `json:"project_id,omitempty"`
	// 一次健康检查请求的超时时间。  建议该值小于delay的值。

	Timeout int32 `json:"timeout"`
	// 健康检查请求协议。 取值：TCP、UDP_CONNECT、HTTP、HTTPS。 使用说明： - 若pool的protocol为QUIC，则type只能是UDP_CONNECT。 - 若pool的protocol为UDP，则type只能UDP_CONNECT。 - 若pool的protocol为TCP，则type可以是TCP、HTTP、HTTPS。 - 若pool的protocol为HTTP，则type可以是TCP、HTTP、HTTPS。 - 若pool的protocol为HTTPS，则type可以是TCP、HTTP、HTTPS。

	Type string `json:"type"`
	// 健康检查请求的请求路径。以\"/\"开头，默认为\"/\"。  使用说明： - 仅当type为HTTP时生效。

	UrlPath *string `json:"url_path,omitempty"`
}

func (o CreateHealthMonitorOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHealthMonitorOption struct{}"
	}

	return strings.Join([]string{"CreateHealthMonitorOption", string(data)}, " ")
}
