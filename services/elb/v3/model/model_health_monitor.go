package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HealthMonitor 健康检查对象
type HealthMonitor struct {

	// 健康检查的管理状态。  取值范围： - true：表示开启健康检查。 - false表示关闭健康检查。  默认取值：true。
	AdminStateUp bool `json:"admin_state_up"`

	// 健康检查间隔。取值：1-50s。
	Delay int32 `json:"delay"`

	// 发送健康检查请求的域名。  取值：以数字或字母开头，只能包含数字、字母、’-’、’.’。 默认为空，表示使用负载均衡器的vip作为http请求的目的地址。  使用说明：当type为HTTP/HTTPS时生效。
	DomainName string `json:"domain_name"`

	// 期望响应状态码。  取值： - 单值：单个返回码，例如200。 - 列表：多个特定返回码，例如200，202。 - 区间：一个返回码区间，例如200-204。   默认值：若健康检查type为gRPC，则默认值为0,；其他为200。  仅支持HTTP/HTTPS/gRPC设置该字段，其他协议设置不会生效。
	ExpectedCodes string `json:"expected_codes"`

	// HTTP请求方法。  取值：GET、HEAD、POST，默认GET。  使用说明：当type为HTTP/HTTPS时生效。
	HttpMethod string `json:"http_method"`

	// 健康检查ID
	Id string `json:"id"`

	// 健康检查连续成功多少次后，将后端服务器的健康检查状态由OFFLINE判定为ONLINE。取值范围：1-10。
	MaxRetries int32 `json:"max_retries"`

	// 健康检查连续失败多少次后，将后端服务器的健康检查状态由ONLINE判定为OFFLINE。取值范围：1-10，默认3。
	MaxRetriesDown int32 `json:"max_retries_down"`

	// 健康检查端口号。取值：1-65535，默认为空，表示使用后端云服务器端口号。[当pool协议为IP时，monitor_port必须指定为非0值。](tag:hws_eu)
	MonitorPort int32 `json:"monitor_port"`

	// 健康检查名称。
	Name string `json:"name"`

	// 健康检查所在的后端云服务器组ID列表。实际只会有一个后端云服务器组ID。
	Pools []PoolRef `json:"pools"`

	// 健康检查所在的项目ID。
	ProjectId string `json:"project_id"`

	// 一次健康检查请求的超时时间。  建议该值小于delay的值。
	Timeout int32 `json:"timeout"`

	// 健康检查请求协议。  取值：TCP、UDP_CONNECT、HTTP、HTTPS、TLS和gRPC。  使用说明： - 若pool的protocol为QUIC，则type只能是UDP_CONNECT。 - 若pool的protocol为UDP，则type只能UDP_CONNECT。 - 若pool的protocol为TCP，则type可以是TCP、HTTP、HTTPS。 [- 若pool的protocol为IP，则type可以是TCP、HTTP、HTTPS。](tag:hws_eu) - 若pool的protocol为HTTP，则type可以是TCP、HTTP、HTTPS、TLS、gRPC。 - 若pool的protocol为HTTPS，则type可以是TCP、HTTP、HTTPS、TLS、gRPC。 - 若pool的protocol为gRPC，则type可以是TCP、HTTP、HTTPS、TLS、gRPC。 - 若pool的protocol为TLS，则type可以是TCP、HTTP、HTTPS、TLS、gRPC。  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt)  [荷兰region不支持QUIC。](tag:dt,dt_test)
	Type string `json:"type"`

	// 健康检查请求的请求路径。以\"/\"开头，默认为\"/\"。  支持使用字母、数字和短划线（-）、正斜线（/）、半角句号（.）、百分号（%）、半角问号（?）、井号（#）和and（&）以及扩展字符集_;~!()*[]@$^:',+   使用说明：当type为HTTP/HTTPS时生效。
	UrlPath string `json:"url_path"`

	// 创建时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。  [注意：独享型实例的历史数据以及共享型实例下的资源，不返回该字段。 ](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,fcs,dt,hk_tm)
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。  [注意：独享型实例的历史数据以及共享型实例下的资源，不返回该字段。 ](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,fcs,dt,hk_tm)
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o HealthMonitor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthMonitor struct{}"
	}

	return strings.Join([]string{"HealthMonitor", string(data)}, " ")
}
