package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListenerProtocol 监听的网络传输协议类型。
type ListenerProtocol struct {
	value string
}

type ListenerProtocolEnum struct {
	TCP ListenerProtocol
	UDP ListenerProtocol
}

func GetListenerProtocolEnum() ListenerProtocolEnum {
	return ListenerProtocolEnum{
		TCP: ListenerProtocol{
			value: "TCP",
		},
		UDP: ListenerProtocol{
			value: "UDP",
		},
	}
}

func (c ListenerProtocol) Value() string {
	return c.value
}

func (c ListenerProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListenerProtocol) UnmarshalJSON(b []byte) error {
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
