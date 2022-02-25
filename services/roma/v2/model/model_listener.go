package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 监听器
type Listener struct {
	// 监听器名称

	Name *string `json:"name,omitempty"`
	// 监听器对外提供服务端口

	Port *int32 `json:"port,omitempty"`

	Backend *Backend `json:"backend,omitempty"`
	// 创建负载均衡器的IP协议类型

	IpVersion *ListenerIpVersion `json:"ip_version,omitempty"`
}

func (o Listener) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Listener struct{}"
	}

	return strings.Join([]string{"Listener", string(data)}, " ")
}

type ListenerIpVersion struct {
	value int32
}

type ListenerIpVersionEnum struct {
	E_4 ListenerIpVersion
	E_6 ListenerIpVersion
}

func GetListenerIpVersionEnum() ListenerIpVersionEnum {
	return ListenerIpVersionEnum{
		E_4: ListenerIpVersion{
			value: 4,
		}, E_6: ListenerIpVersion{
			value: 6,
		},
	}
}

func (c ListenerIpVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListenerIpVersion) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
