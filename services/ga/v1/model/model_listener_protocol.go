package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 监听的网络传输协议类型。
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
