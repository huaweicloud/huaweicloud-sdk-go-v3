package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 创建后端云服务器组的请求体
type CreatePoolReq struct {
	// 后端云服务器组的后端协议。取值：UDP、TCP、HTTP。当指定istener_id创建后端云服务器组时，后端云服务器组的protocol和它关联的监听器的protocol有如下关系：监听器的protocol为TCP时，后端云服务器组的protocol必须为TCP。监听器的protocol为UDP时，后端云服务器组的protocol必须为UDP。监听器的protocol为HTTP或TERMINATED_HTTPS时，后端云服务器组的protocol必须为HTTP。

	Protocol CreatePoolReqProtocol `json:"protocol"`
	// 后端云服务器组的负载均衡算法，取值：ROUND_ROBIN：加权轮询算法；LEAST_CONNECTIONS：加权最少连接算法；SOURCE_IP：源IP算法；当该字段的取值为SOURCE_IP时，后端云服务器组绑定的后端云服务器的weight字段无效。

	LbAlgorithm string `json:"lb_algorithm"`
	// 后端云服务器组关联的负载均衡器ID。listener_id和loadbalancer_id中至少指定一个。

	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`
	// 后端云服务器组关联的监听器的ID。listener_id和loadbalancer_id中至少指定一个。

	ListenerId *string `json:"listener_id,omitempty"`
	// 后端云服务器组所在的项目ID。

	TenantId *string `json:"tenant_id,omitempty"`
	// 后端云服务器组的名称。

	Name *string `json:"name,omitempty"`
	// 后端云服务器组的描述信息

	Description *string `json:"description,omitempty"`
	// 后端云服务器组的管理状态。只支持设定为true，该字段的值无实际意义。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	SessionPersistence *SessionPersistence `json:"session_persistence,omitempty"`
}

func (o CreatePoolReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePoolReq struct{}"
	}

	return strings.Join([]string{"CreatePoolReq", string(data)}, " ")
}

type CreatePoolReqProtocol struct {
	value string
}

type CreatePoolReqProtocolEnum struct {
	UDP  CreatePoolReqProtocol
	TCP  CreatePoolReqProtocol
	HTTP CreatePoolReqProtocol
}

func GetCreatePoolReqProtocolEnum() CreatePoolReqProtocolEnum {
	return CreatePoolReqProtocolEnum{
		UDP: CreatePoolReqProtocol{
			value: "UDP",
		},
		TCP: CreatePoolReqProtocol{
			value: "TCP",
		},
		HTTP: CreatePoolReqProtocol{
			value: "HTTP",
		},
	}
}

func (c CreatePoolReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePoolReqProtocol) UnmarshalJSON(b []byte) error {
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
