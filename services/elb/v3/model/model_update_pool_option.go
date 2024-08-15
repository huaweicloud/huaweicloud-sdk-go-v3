package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdatePoolOption 更新后端服务器组请求参数。
type UpdatePoolOption struct {

	// 后端云服务器组的管理状态，只支持更新为true。  [不支持该字段，请勿使用。](tag:dt,dt_test,hcso_dt)
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 后端云服务器组的描述信息。
	Description *string `json:"description,omitempty"`

	// 后端云服务器组的负载均衡算法。  取值： - ROUND_ROBIN：加权轮询算法。 - LEAST_CONNECTIONS：加权最少连接算法。 - SOURCE_IP：源IP算法。 - QUIC_CID：连接ID算法。 [- 2_TUPLE_HASH：二元组hash算法，仅IP类型的pool支持。 - 3_TUPLE_HASH：三元组hash算法，仅IP类型的pool支持。 - 5_TUPLE_HASH：五元组hash算法，仅IP类型的pool支持。 - IP型pool不指定该字段时，默认设置为5_TUPLE_HASH。](tag:hws_eu)  使用说明： - 当该字段的取值为SOURCE_IP时，后端云服务器组绑定的后端云服务器的weight字段无效。 - 只有pool的protocol为QUIC时，才支持QUIC_CID算法。  [不支持QUIC_CID。](tag:tm,hws_eu,g42,hk_g42,hcso_dt)  [荷兰region不支持QUIC_CID。](tag:dt,dt_test)
	LbAlgorithm *string `json:"lb_algorithm,omitempty"`

	// 后端云服务器组的名称。
	Name *string `json:"name,omitempty"`

	SessionPersistence *UpdatePoolSessionPersistenceOption `json:"session_persistence,omitempty"`

	SlowStart *UpdatePoolSlowStartOption `json:"slow_start,omitempty"`

	// 是否开启删除保护。  取值：false不开启，true开启。  > 退场时需要先关闭所有资源的删除保护开关。  [不支持该字段，请勿使用。](tag:hws_eu,g42,hk_g42)  [荷兰region不支持该字段，请勿使用。](tag:dt,dt_test)
	MemberDeletionProtectionEnable *bool `json:"member_deletion_protection_enable,omitempty"`

	// 后端云服务器组关联的虚拟私有云的ID。 [- pool的protocol为IP时，必须指定vpc_id，且与LB的vpc_id相同。](tag:hws_eu)  只有vpc_id为空时允许更新。
	VpcId *string `json:"vpc_id,omitempty"`

	// 后端服务器组的类型。  取值： - instance：允许任意类型的后端，type指定为该类型时，vpc_id是必选字段。 - ip：只能添加跨VPC后端，type指定为该类型时，vpc_id不允许指定。[pool的protocol为IP时，type不允许设置为ip。](tag:hws_eu)] - 空字符串（\"\"）：允许任意类型的后端  使用说明： - 只有type为空时允许更新，不允许从非空更新为空。
	Type *string `json:"type,omitempty"`

	// 修改保护状态, 取值： - nonProtection: 不保护 - consoleProtection: 控制台修改保护
	ProtectionStatus *UpdatePoolOptionProtectionStatus `json:"protection_status,omitempty"`

	// 设置保护的原因 >仅当protection_status为consoleProtection时有效。
	ProtectionReason *string `json:"protection_reason,omitempty"`

	// 后端是否开启端口透传。开启后，后端服务器端口与前端监听器端口保持一致。关闭后，请求会转发给后端服务器protocol_port字段指定端口。取值：false不开启，true开启。  使用说明： - 仅QUIC,TCP,UDP的pool支持。
	AnyPortEnable *bool `json:"any_port_enable,omitempty"`

	ConnectionDrain *ConnectionDrain `json:"connection_drain,omitempty"`

	PoolHealth *PoolHealth `json:"pool_health,omitempty"`
}

func (o UpdatePoolOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePoolOption struct{}"
	}

	return strings.Join([]string{"UpdatePoolOption", string(data)}, " ")
}

type UpdatePoolOptionProtectionStatus struct {
	value string
}

type UpdatePoolOptionProtectionStatusEnum struct {
	NON_PROTECTION     UpdatePoolOptionProtectionStatus
	CONSOLE_PROTECTION UpdatePoolOptionProtectionStatus
}

func GetUpdatePoolOptionProtectionStatusEnum() UpdatePoolOptionProtectionStatusEnum {
	return UpdatePoolOptionProtectionStatusEnum{
		NON_PROTECTION: UpdatePoolOptionProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: UpdatePoolOptionProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c UpdatePoolOptionProtectionStatus) Value() string {
	return c.value
}

func (c UpdatePoolOptionProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePoolOptionProtectionStatus) UnmarshalJSON(b []byte) error {
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
