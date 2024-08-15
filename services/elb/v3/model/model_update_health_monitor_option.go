package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateHealthMonitorOption 更新健康检查请求参数。
type UpdateHealthMonitorOption struct {

	// 参数解释：健康检查的管理状态。  取值范围： - true：表示开启健康检查。 - false表示关闭健康检查。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 健康检查间隔。健康检查间隔。取值：1-50s。
	Delay *int32 `json:"delay,omitempty"`

	// 发送健康检查请求的域名。  取值：以数字或字母开头，只能包含数字、字母、’-’、’.’。 不能传空，但可传null或不传，表示使用负载均衡器的vip作为http请求的目的地址。  使用说明：当type为HTTP/HTTPS时生效。
	DomainName *string `json:"domain_name,omitempty"`

	// 期望响应状态码。  取值： - 单值：单个返回码，例如200。 - 列表：多个特定返回码，例如200，202。 - 区间：一个返回码区间，例如200-204。  仅支持HTTP/HTTPS/gRPC设置该字段，其他协议设置不会生效。
	ExpectedCodes *string `json:"expected_codes,omitempty"`

	// HTTP请求方法。  取值：GET、HEAD、POST，默认GET。  使用说明：当type为HTTP/HTTPS时生效。
	HttpMethod *UpdateHealthMonitorOptionHttpMethod `json:"http_method,omitempty"`

	// 健康检查连续成功多少次后，将后端服务器的健康检查状态由OFFLINE判定为ONLINE。取值范围：1-10。
	MaxRetries *int32 `json:"max_retries,omitempty"`

	// 健康检查连续失败多少次后，将后端服务器的健康检查状态由ONLINE判定为OFFLINE。取值范围：1-10。
	MaxRetriesDown *int32 `json:"max_retries_down,omitempty"`

	// 健康检查端口号。取值：1-65535，不可传入空，但可传入null，表示使用后端云服务器端口号。[当pool协议为IP时，monitor_port必须指定为非0值。](tag:hws_eu)
	MonitorPort *int32 `json:"monitor_port,omitempty"`

	// 健康检查名称。
	Name *string `json:"name,omitempty"`

	// 一次健康检查请求的超时时间。  建议该值小于delay的值。
	Timeout *int32 `json:"timeout,omitempty"`

	// 健康检查请求的请求路径。以\"/\"开头，默认为\"/\"。  支持使用字母、数字和短划线（-）、正斜线（/）、半角句号（.）、百分号（%）、半角问号（?）、井号（#）和and（&）以及扩展字符集_;~!()*[]@$^:',+  使用说明：当type为HTTP/HTTPS时生效。
	UrlPath *string `json:"url_path,omitempty"`

	// 健康检查请求协议。  取值：TCP、UDP_CONNECT、HTTP、HTTPS、TLS和gRPC。  使用说明： - 若pool的protocol为QUIC，则type只能是UDP_CONNECT。 - 若pool的protocol为UDP，则type只能UDP_CONNECT。 - 若pool的protocol为TCP，则type可以是TCP、HTTP、HTTPS。 [- 若pool的protocol为IP，则type可以是TCP、HTTP、HTTPS。](tag:hws_eu) - 若pool的protocol为HTTP，则type可以是TCP、HTTP、HTTPS、TLS、gRPC。 - 若pool的protocol为HTTPS，则type可以是TCP、HTTP、HTTPS、TLS、gRPC。 - 若pool的protocol为gRPC，则type可以是TCP、HTTP、HTTPS、TLS、gRPC。 - 若pool的protocol为TLS，则type可以是TCP、HTTP、HTTPS、TLS、gRPC。  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt)  [荷兰region不支持QUIC。](tag:dt,dt_test)
	Type *string `json:"type,omitempty"`
}

func (o UpdateHealthMonitorOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHealthMonitorOption struct{}"
	}

	return strings.Join([]string{"UpdateHealthMonitorOption", string(data)}, " ")
}

type UpdateHealthMonitorOptionHttpMethod struct {
	value string
}

type UpdateHealthMonitorOptionHttpMethodEnum struct {
	GET  UpdateHealthMonitorOptionHttpMethod
	HEAD UpdateHealthMonitorOptionHttpMethod
	POST UpdateHealthMonitorOptionHttpMethod
}

func GetUpdateHealthMonitorOptionHttpMethodEnum() UpdateHealthMonitorOptionHttpMethodEnum {
	return UpdateHealthMonitorOptionHttpMethodEnum{
		GET: UpdateHealthMonitorOptionHttpMethod{
			value: "GET",
		},
		HEAD: UpdateHealthMonitorOptionHttpMethod{
			value: "HEAD",
		},
		POST: UpdateHealthMonitorOptionHttpMethod{
			value: "POST",
		},
	}
}

func (c UpdateHealthMonitorOptionHttpMethod) Value() string {
	return c.value
}

func (c UpdateHealthMonitorOptionHttpMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateHealthMonitorOptionHttpMethod) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
