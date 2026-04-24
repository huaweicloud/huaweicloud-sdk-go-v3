package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BindDNatRequestBody struct {

	// **参数解释**: 需要绑定或者解绑的节点ID。 **约束限制**: 分布式仅支持CN节点，集中式不支持日志节点。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	NodeId string `json:"node_id"`

	// **参数解释**: 弹性公网ID。 **约束限制**: action类型为BIND时必选。 一个弹性公网IP只能绑定到一个NAT网关。 **取值范围**: UUID格式。 **默认取值**: 不涉及。
	PublicIpId *string `json:"public_ip_id,omitempty"`

	// **参数解释**: 公网NAT网关的ID。 **约束限制**: action类型为BIND时必选。 NAT网关的虚拟私有云和子网需要和GaussDB数据库实例的虚拟私有云和子网保持一致。 **取值范围**: UUID格式。 **默认取值**: 不涉及。
	NatGatewayId *string `json:"nat_gateway_id,omitempty"`

	// **参数解释**: 对外提供服务的端口号，可通过弹性公网IP加该端口号的方式连接数据库实例。 **约束限制**: action类型为BIND时必选。 **取值范围**: 0~65535。 **默认取值**: 不涉及。
	ExternalServicePort *int32 `json:"external_service_port,omitempty"`

	// **参数解释**: 操作标识。 **约束限制**: 不涉及。 **取值范围**: BIND，表示绑定NAT网关。 UNBIND，表示解绑NAT网关。 **默认取值**: 不涉及。
	Action BindDNatRequestBodyAction `json:"action"`
}

func (o BindDNatRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindDNatRequestBody struct{}"
	}

	return strings.Join([]string{"BindDNatRequestBody", string(data)}, " ")
}

type BindDNatRequestBodyAction struct {
	value string
}

type BindDNatRequestBodyActionEnum struct {
	BIND   BindDNatRequestBodyAction
	UNBIND BindDNatRequestBodyAction
}

func GetBindDNatRequestBodyActionEnum() BindDNatRequestBodyActionEnum {
	return BindDNatRequestBodyActionEnum{
		BIND: BindDNatRequestBodyAction{
			value: "BIND",
		},
		UNBIND: BindDNatRequestBodyAction{
			value: "UNBIND",
		},
	}
}

func (c BindDNatRequestBodyAction) Value() string {
	return c.value
}

func (c BindDNatRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BindDNatRequestBodyAction) UnmarshalJSON(b []byte) error {
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
