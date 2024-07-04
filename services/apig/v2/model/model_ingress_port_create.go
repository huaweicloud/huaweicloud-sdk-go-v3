package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// IngressPortCreate 实例自定义入方向端口创建信息。
type IngressPortCreate struct {

	// 实例自定义入方向端口协议。 - HTTP：实例自定义入方向端口使用HTTP协议。 - HTTPS：实例自定义入方向端口使用HTTPS协议。
	Protocol IngressPortCreateProtocol `json:"protocol"`

	// 实例自定义入方向端口，支持的端口范围为1024~49151。
	IngressPort int32 `json:"ingress_port"`
}

func (o IngressPortCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IngressPortCreate struct{}"
	}

	return strings.Join([]string{"IngressPortCreate", string(data)}, " ")
}

type IngressPortCreateProtocol struct {
	value string
}

type IngressPortCreateProtocolEnum struct {
	HTTP  IngressPortCreateProtocol
	HTTPS IngressPortCreateProtocol
}

func GetIngressPortCreateProtocolEnum() IngressPortCreateProtocolEnum {
	return IngressPortCreateProtocolEnum{
		HTTP: IngressPortCreateProtocol{
			value: "HTTP",
		},
		HTTPS: IngressPortCreateProtocol{
			value: "HTTPS",
		},
	}
}

func (c IngressPortCreateProtocol) Value() string {
	return c.value
}

func (c IngressPortCreateProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IngressPortCreateProtocol) UnmarshalJSON(b []byte) error {
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
