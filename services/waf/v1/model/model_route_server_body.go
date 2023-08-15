package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RouteServerBody 防护域名源站服务器信息
type RouteServerBody struct {

	// WAF转发客户端请求到防护域名源站服务器的协议
	BackProtocol *RouteServerBodyBackProtocol `json:"back_protocol,omitempty"`

	// 客户端访问的源站服务器的IP地址
	Address *string `json:"address,omitempty"`

	// WAF转发客户端请求到源站服务的业务端口
	Port *int32 `json:"port,omitempty"`
}

func (o RouteServerBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RouteServerBody struct{}"
	}

	return strings.Join([]string{"RouteServerBody", string(data)}, " ")
}

type RouteServerBodyBackProtocol struct {
	value string
}

type RouteServerBodyBackProtocolEnum struct {
	HTTP  RouteServerBodyBackProtocol
	HTTPS RouteServerBodyBackProtocol
}

func GetRouteServerBodyBackProtocolEnum() RouteServerBodyBackProtocolEnum {
	return RouteServerBodyBackProtocolEnum{
		HTTP: RouteServerBodyBackProtocol{
			value: "HTTP",
		},
		HTTPS: RouteServerBodyBackProtocol{
			value: "HTTPS",
		},
	}
}

func (c RouteServerBodyBackProtocol) Value() string {
	return c.value
}

func (c RouteServerBodyBackProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RouteServerBodyBackProtocol) UnmarshalJSON(b []byte) error {
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
