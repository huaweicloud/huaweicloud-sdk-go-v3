package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateElbClusterPortRequestBody struct {

	// 集群id
	ClusterId string `json:"cluster_id"`

	// elb id。端口映射将会在该elb中创建
	ElbId string `json:"elb_id"`

	// 新增的端口的模式。除了VPN外，其他类型的服务只支持 PROXY
	Mode CreateElbClusterPortRequestBodyMode `json:"mode"`

	// 将在elb中为代理通道创建的监听器的端口
	LbListenerPort int32 `json:"lb_listener_port"`

	// 将在elb中为代理通道创建的后端服务组中后端服务器的端口。无默认值。
	ServerPort int32 `json:"server_port"`

	// 将在elb中给VPN隧道创建的监听器的端口。当mode=TUNNEL时，必填
	TunnelLbListenerPort *int32 `json:"tunnel_lb_listener_port,omitempty"`

	// 将在elb中给VPN隧道创建的后端服务组中后端服务器的端口。 当mode=TUNNEL时，必填，有默认值:20706。
	TunnelServerPort *int32 `json:"tunnel_server_port,omitempty"`
}

func (o CreateElbClusterPortRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateElbClusterPortRequestBody struct{}"
	}

	return strings.Join([]string{"CreateElbClusterPortRequestBody", string(data)}, " ")
}

type CreateElbClusterPortRequestBodyMode struct {
	value string
}

type CreateElbClusterPortRequestBodyModeEnum struct {
	TUNNEL CreateElbClusterPortRequestBodyMode
	PROXY  CreateElbClusterPortRequestBodyMode
}

func GetCreateElbClusterPortRequestBodyModeEnum() CreateElbClusterPortRequestBodyModeEnum {
	return CreateElbClusterPortRequestBodyModeEnum{
		TUNNEL: CreateElbClusterPortRequestBodyMode{
			value: "TUNNEL",
		},
		PROXY: CreateElbClusterPortRequestBodyMode{
			value: "PROXY",
		},
	}
}

func (c CreateElbClusterPortRequestBodyMode) Value() string {
	return c.value
}

func (c CreateElbClusterPortRequestBodyMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateElbClusterPortRequestBodyMode) UnmarshalJSON(b []byte) error {
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
