package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// HealthCheck 健康检查配置
type HealthCheck struct {

	// 当前LVS只支持TCP_CHECK
	Protocol *HealthCheckProtocol `json:"protocol,omitempty"`

	// 健康检查端口
	Port *int32 `json:"port,omitempty"`

	// 判定后端不健康的阈值，连续探测失败次数
	Unhealthy *int32 `json:"unhealthy,omitempty"`

	// 探测后端的超时时间，单位秒
	Timeout *int32 `json:"timeout,omitempty"`

	// 探测后端的间隔时间，单位秒
	Interval *int32 `json:"interval,omitempty"`
}

func (o HealthCheck) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthCheck struct{}"
	}

	return strings.Join([]string{"HealthCheck", string(data)}, " ")
}

type HealthCheckProtocol struct {
	value string
}

type HealthCheckProtocolEnum struct {
	TCP_CHECK HealthCheckProtocol
	HTTP_GET  HealthCheckProtocol
	SSL_GET   HealthCheckProtocol
}

func GetHealthCheckProtocolEnum() HealthCheckProtocolEnum {
	return HealthCheckProtocolEnum{
		TCP_CHECK: HealthCheckProtocol{
			value: "TCP_CHECK",
		},
		HTTP_GET: HealthCheckProtocol{
			value: "HTTP_GET",
		},
		SSL_GET: HealthCheckProtocol{
			value: "SSL_GET",
		},
	}
}

func (c HealthCheckProtocol) Value() string {
	return c.value
}

func (c HealthCheckProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HealthCheckProtocol) UnmarshalJSON(b []byte) error {
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
