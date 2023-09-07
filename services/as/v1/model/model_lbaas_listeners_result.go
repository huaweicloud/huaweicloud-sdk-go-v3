package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LbaasListenersResult 增强型负载均衡器信息
type LbaasListenersResult struct {

	// 监听器ID
	ListenerId *string `json:"listener_id,omitempty"`

	// 后端云服务器组ID
	PoolId *string `json:"pool_id,omitempty"`

	// 后端协议号，指后端云服务器监听的端口，取值范围[1,65535]
	ProtocolPort *int32 `json:"protocol_port,omitempty"`

	// 权重，指后端云服务器经分发得到的请求数量比例，取值范围[0,1]，默认为1。
	Weight *int32 `json:"weight,omitempty"`

	// 指定ip协议版本
	ProtocolVersion *LbaasListenersResultProtocolVersion `json:"protocol_version,omitempty"`
}

func (o LbaasListenersResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LbaasListenersResult struct{}"
	}

	return strings.Join([]string{"LbaasListenersResult", string(data)}, " ")
}

type LbaasListenersResultProtocolVersion struct {
	value string
}

type LbaasListenersResultProtocolVersionEnum struct {
	IPV4 LbaasListenersResultProtocolVersion
	IPV6 LbaasListenersResultProtocolVersion
}

func GetLbaasListenersResultProtocolVersionEnum() LbaasListenersResultProtocolVersionEnum {
	return LbaasListenersResultProtocolVersionEnum{
		IPV4: LbaasListenersResultProtocolVersion{
			value: "IPV4",
		},
		IPV6: LbaasListenersResultProtocolVersion{
			value: "IPV6",
		},
	}
}

func (c LbaasListenersResultProtocolVersion) Value() string {
	return c.value
}

func (c LbaasListenersResultProtocolVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LbaasListenersResultProtocolVersion) UnmarshalJSON(b []byte) error {
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
