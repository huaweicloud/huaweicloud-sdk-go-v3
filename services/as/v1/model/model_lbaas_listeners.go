package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LbaasListeners 增强型负载均衡器
type LbaasListeners struct {

	// 后端云服务器组ID
	PoolId string `json:"pool_id"`

	// 后端协议号，指后端云服务器监听的端口，取值范围[1,65535]
	ProtocolPort int32 `json:"protocol_port"`

	// 权重，指后端云服务器经分发得到的请求数量的比例，取值范围[0, 100]。
	Weight int32 `json:"weight"`

	// 指定ip协议版本
	ProtocolVersion *LbaasListenersProtocolVersion `json:"protocol_version,omitempty"`
}

func (o LbaasListeners) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LbaasListeners struct{}"
	}

	return strings.Join([]string{"LbaasListeners", string(data)}, " ")
}

type LbaasListenersProtocolVersion struct {
	value string
}

type LbaasListenersProtocolVersionEnum struct {
	IPV4 LbaasListenersProtocolVersion
	IPV6 LbaasListenersProtocolVersion
}

func GetLbaasListenersProtocolVersionEnum() LbaasListenersProtocolVersionEnum {
	return LbaasListenersProtocolVersionEnum{
		IPV4: LbaasListenersProtocolVersion{
			value: "IPV4",
		},
		IPV6: LbaasListenersProtocolVersion{
			value: "IPV6",
		},
	}
}

func (c LbaasListenersProtocolVersion) Value() string {
	return c.value
}

func (c LbaasListenersProtocolVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LbaasListenersProtocolVersion) UnmarshalJSON(b []byte) error {
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
