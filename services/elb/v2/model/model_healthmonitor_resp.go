package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 健康检查响应体
type HealthmonitorResp struct {
	// 健康检查ID

	Id string `json:"id"`
	// 健康检查所在的项目ID。

	ProjectId string `json:"project_id"`
	// 健康检查所在的项目ID。

	TenantId string `json:"tenant_id"`
	// 健康检查名称。

	Name string `json:"name"`
	// 健康检查的管理状态；该字段虽然支持创建、更新，但实际取值决定于后端云服务器对应的弹性云服务器是否存在。该字段虽然支持创建、更新，但实际取值决定于member对应的弹性云服务器是否存在。若存在，该值为true，否则，该值为false。

	AdminStateUp bool `json:"admin_state_up"`
	// 健康检查端口号。默认为空，表示使用后端云服务器组的端口。

	MonitorPort int32 `json:"monitor_port"`
	// 健康检查的超时时间。建议该值小于delay的值。

	Timeout int32 `json:"timeout"`
	// 健康检查类型

	Type HealthmonitorRespType `json:"type"`
	// 期望HTTP响应状态码，指定下列值：单值，例如200；列表，例如200，202；区间，例如200-204。仅当type为HTTP时生效。该字段为预留字段，暂未启用。

	ExpectedCodes string `json:"expected_codes"`
	// 功能说明：健康检查测试member健康状态时，发送的http请求的域名。仅当type为HTTP时生效。使用说明：默认为空，表示使用负载均衡器的vip作为http请求的目的地址。以数字或字母开头，只能包含数字、字母、’-’、’.’。

	DomainName string `json:"domain_name"`
	// HTTP方法，可以为GET、HEAD、POST、PUT、DELETE、TRACE、OPTIONS、CONNECT、PATCH。仅当type为HTTP时生效。该字段为预留字段，暂未启用。

	UrlPath string `json:"url_path"`
	// HTTP方法，可以为GET、HEAD、POST、PUT、DELETE、TRACE、OPTIONS、CONNECT、PATCH。仅当type为HTTP时生效。该字段为预留字段，暂未启用。

	HttpMethod string `json:"http_method"`
	// 健康检查间隔，单位秒

	Delay int32 `json:"delay"`
	// 最大重试次数

	MaxRetries int32 `json:"max_retries"`
	// 健康检查关联的后端云服务器组列表

	Pools []ResourceList `json:"pools"`
}

func (o HealthmonitorResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthmonitorResp struct{}"
	}

	return strings.Join([]string{"HealthmonitorResp", string(data)}, " ")
}

type HealthmonitorRespType struct {
	value string
}

type HealthmonitorRespTypeEnum struct {
	TCP         HealthmonitorRespType
	UDP_CONNECT HealthmonitorRespType
	HTTP        HealthmonitorRespType
}

func GetHealthmonitorRespTypeEnum() HealthmonitorRespTypeEnum {
	return HealthmonitorRespTypeEnum{
		TCP: HealthmonitorRespType{
			value: "TCP",
		},
		UDP_CONNECT: HealthmonitorRespType{
			value: "UDP_CONNECT",
		},
		HTTP: HealthmonitorRespType{
			value: "HTTP",
		},
	}
}

func (c HealthmonitorRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HealthmonitorRespType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
