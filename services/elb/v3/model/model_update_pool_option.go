package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdatePoolOption struct {

	// **参数解释**：后端服务器组的管理状态。  **约束限制**：只支持更新为true。  **取值范围**：true 启用。  **默认取值**：不涉及  [不支持该字段，请勿使用。](tag:dt,hcso_dt)
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// **参数解释**：后端服务器组的描述信息。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**：后端服务器组的负载均衡算法。  **约束限制**： - 当该字段的取值为SOURCE_IP或QUIC_CID时，后端服务器组绑定的后端服务器的weight字段无效。 - 只有pool的protocol为QUIC时，才支持QUIC_CID算法。  **取值范围**： - ROUND_ROBIN：加权轮询算法。 - LEAST_CONNECTIONS：加权最少连接算法。 - SOURCE_IP：源IP算法。 - QUIC_CID：连接ID算法。 [- 2_TUPLE_HASH：二元组hash算法，仅IP类型的pool支持。 - 3_TUPLE_HASH：三元组hash算法，仅IP类型的pool支持。 - 5_TUPLE_HASH：五元组hash算法，仅IP类型的pool支持。 - IP型pool不指定该字段时，默认设置为5_TUPLE_HASH。](tag:hws_eu)  **默认取值**：不涉及  [不支持QUIC_CID。](tag:tm,hws_eu,g42,hk_g42,hcso_dt)  [荷兰region不支持QUIC_CID。](tag:dt)
	LbAlgorithm *string `json:"lb_algorithm,omitempty"`

	// **参数解释**：后端服务器组的名称。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Name *string `json:"name,omitempty"`

	SessionPersistence *UpdatePoolSessionPersistenceOption `json:"session_persistence,omitempty"`

	SlowStart *UpdatePoolSlowStartOption `json:"slow_start,omitempty"`

	// **参数解释**：是否开启后端服务器移除保护。开关开启后，不允许从该ELB后端服务器组下移除后端服务器。  **约束限制**： - 开关开启后，移除member会报错拦截，涉及如下API:   + 级联删除负载均衡器（DELETE /v3/{project_id}/elb/loadbalancers/{loadbalancer_id}/force-elb）   + 级联删除负载均衡器及关联EIP（POST /v3/{project_id}/elb/loadbalancers/{loadbalancer_id}/delete-cascade）   + 级联删除监听器（DELETE /v3/{project_id}/elb/listeners/{listener_id}/force）   + 级联删除后端服务器组（DELETE /v3/{project_id}/elb/pools/{pool_id}/delete-cascade）   + 删除后端服务器（DELETE /v3/{project_id}/elb/pools/{pool_id}/members/{member_id}）   + 批量删除后端服务器（POST /v3/{project_id}/elb/pools/{pool_id}/members/batch-delete）   **取值范围**：false 不开启，true 开启。  **默认取值**：false  > 退场时需要先关闭所有资源的删除保护开关。  [不支持该字段，请勿使用。](tag:hws_eu,g42,hk_g42)  [荷兰region不支持该字段，请勿使用。](tag:dt)
	MemberDeletionProtectionEnable *bool `json:"member_deletion_protection_enable,omitempty"`

	// **参数解释**：后端服务器组关联的虚拟私有云的ID。  **约束限制**：只有vpc_id为空时允许更新。  **取值范围**：不涉及  **默认取值**：不涉及
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释**：后端服务器组的类型。  **约束限制**： - 只有type为空时允许更新，不允许从非空更新为空。 - instance：允许任意类型的后端，type指定为该类型时，vpc_id是必选字段。 - ip：只能添加IP类型后端，type指定为该类型时，vpc_id不允许指定。[pool的protocol为IP时，type不允许设置为ip。](tag:hws_eu)] - 空字符串（\"\"）：允许任意类型的后端  **取值范围**：不涉及  **默认取值**：不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释**：修改保护状态。  **约束限制**：不涉及  **取值范围**： - nonProtection: 不保护 - consoleProtection: 控制台修改保护  **默认取值**：不涉及
	ProtectionStatus *UpdatePoolOptionProtectionStatus `json:"protection_status,omitempty"`

	// **参数解释**：设置保护的原因。作为protection_status的转态设置的原因。  **约束限制**：仅当protection_status为consoleProtection时有效。  **取值范围**：除'<'和'>'外通用Unicode字符集字符，最大255个字符。  **默认取值**：不涉及
	ProtectionReason *string `json:"protection_reason,omitempty"`

	// **参数解释**：后端是否开启全端口转发。开启后，后端服务器端口与前端监听器端口保持一致。关闭后，请求会转发给后端服务器protocol_port字段指定端口。  **约束限制**：仅QUIC,TCP,UDP的pool支持。  **取值范围**：false 不开启，true 开启。  **默认取值**：false
	AnyPortEnable *bool `json:"any_port_enable,omitempty"`

	ConnectionDrain *ConnectionDrain `json:"connection_drain,omitempty"`

	PoolHealth *PoolHealth `json:"pool_health,omitempty"`

	QuicCidHashStrategy *QuicCidHashStrategy `json:"quic_cid_hash_strategy,omitempty"`

	AzAffinity *UpdateAzAffinity `json:"az_affinity,omitempty"`
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
