package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 服务开放的端口映射列表
type PortList struct {
	// 终端节点访问的端口。 终端节点提供给用户，作为访问终端节 点服务的端口，范围1-65535。

	ClientPort *int32 `json:"client_port,omitempty"`
	// 终端节点服务的端口。 终端节点服务绑定了后端资源，作为提 供服务的端口，范围1-65535。

	ServerPort *int32 `json:"server_port,omitempty"`
	// 端口映射协议，支持TCP。

	Protocol *PortListProtocol `json:"protocol,omitempty"`
}

func (o PortList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PortList struct{}"
	}

	return strings.Join([]string{"PortList", string(data)}, " ")
}

type PortListProtocol struct {
	value string
}

type PortListProtocolEnum struct {
	TCP PortListProtocol
}

func GetPortListProtocolEnum() PortListProtocolEnum {
	return PortListProtocolEnum{
		TCP: PortListProtocol{
			value: "TCP",
		},
	}
}

func (c PortListProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PortListProtocol) UnmarshalJSON(b []byte) error {
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
