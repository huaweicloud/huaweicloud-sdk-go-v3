package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PoolResp 后端云服务器组响应体
type PoolResp struct {

	// 后端云服务器组的ID
	Id string `json:"id"`

	// 后端云服务器组所在的项目ID。
	ProjectId string `json:"project_id"`

	// 后端云服务器组所在的项目ID。
	TenantId string `json:"tenant_id"`

	// 后端云服务器组的名称。
	Name string `json:"name"`

	// 后端云服务器组的描述信息
	Description string `json:"description"`

	// 后端云服务器组的管理状态。只支持设定为true，该字段的值无实际意义。
	AdminStateUp bool `json:"admin_state_up"`

	// 后端云服务器组绑定的负载均衡器ID的列表。
	Loadbalancers []ResourceList `json:"loadbalancers"`

	// 后端云服务器组关联的监听器ID的列表。
	Listeners []ResourceList `json:"listeners"`

	// 后端云服务器组关联的后端云服务器ID的列表。
	Members []ResourceList `json:"members"`

	// 后端云服务器组关联的健康检查的ID。
	HealthmonitorId string `json:"healthmonitor_id"`

	SessionPersistence *SessionPersistence `json:"session_persistence"`

	// 后端云服务器组的后端协议。
	Protocol PoolRespProtocol `json:"protocol"`

	// 后端云服务器组的负载均衡算法，取值：ROUND_ROBIN：加权轮询算法；LEAST_CONNECTIONS：加权最少连接算法；SOURCE_IP：源IP算法。当该字段的取值为SOURCE_IP时，后端云服务器组绑定的后端云服务器的weight字段无效。
	LbAlgorithm PoolRespLbAlgorithm `json:"lb_algorithm"`
}

func (o PoolResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolResp struct{}"
	}

	return strings.Join([]string{"PoolResp", string(data)}, " ")
}

type PoolRespProtocol struct {
	value string
}

type PoolRespProtocolEnum struct {
	UDP  PoolRespProtocol
	TCP  PoolRespProtocol
	HTTP PoolRespProtocol
}

func GetPoolRespProtocolEnum() PoolRespProtocolEnum {
	return PoolRespProtocolEnum{
		UDP: PoolRespProtocol{
			value: "UDP",
		},
		TCP: PoolRespProtocol{
			value: "TCP",
		},
		HTTP: PoolRespProtocol{
			value: "HTTP",
		},
	}
}

func (c PoolRespProtocol) Value() string {
	return c.value
}

func (c PoolRespProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoolRespProtocol) UnmarshalJSON(b []byte) error {
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

type PoolRespLbAlgorithm struct {
	value string
}

type PoolRespLbAlgorithmEnum struct {
	ROUND_ROBIN       PoolRespLbAlgorithm
	LEAST_CONNECTIONS PoolRespLbAlgorithm
	SOURCE_IP         PoolRespLbAlgorithm
}

func GetPoolRespLbAlgorithmEnum() PoolRespLbAlgorithmEnum {
	return PoolRespLbAlgorithmEnum{
		ROUND_ROBIN: PoolRespLbAlgorithm{
			value: "ROUND_ROBIN",
		},
		LEAST_CONNECTIONS: PoolRespLbAlgorithm{
			value: "LEAST_CONNECTIONS",
		},
		SOURCE_IP: PoolRespLbAlgorithm{
			value: "SOURCE_IP",
		},
	}
}

func (c PoolRespLbAlgorithm) Value() string {
	return c.value
}

func (c PoolRespLbAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoolRespLbAlgorithm) UnmarshalJSON(b []byte) error {
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
