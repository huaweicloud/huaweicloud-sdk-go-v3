package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreatePoolOption 创建主机组请求
type CreatePoolOption struct {

	// 后端云服务器组的管理状态，只支持更新为true。  不支持该字段，请勿使用。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 后端云服务器组的描述信息。
	Description *string `json:"description,omitempty"`

	// 后端云服务器组的负载均衡算法。  取值： - ROUND_ROBIN：加权轮询算法。 - LEAST_CONNECTIONS：加权最少连接算法。 - SOURCE_IP：源IP算法。 - QUIC_CID：连接ID算法。  使用说明： - 当该字段的取值为SOURCE_IP时，后端云服务器组绑定的后端云服务器的weight字段无效。 - 只有pool的protocol为QUIC时，才支持QUIC_CID算法。  [不支持QUIC_CID。](tag:tm,hws_eu,g42,hk_g42,hcso_dt)  [荷兰region不支持QUIC_CID。](tag:dt,dt_test)
	LbAlgorithm string `json:"lb_algorithm"`

	// 后端云服务器组关联的监听器的ID。  使用说明： - listener_id，loadbalancer_id，type至少指定一个。 [- 共享型实例的后端服务器组loadbalancer_id和listener_id至少指定一个。 ](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,fcs,dt,hk_tm)
	ListenerId *string `json:"listener_id,omitempty"`

	// 后端云服务器组关联的负载均衡器ID。  使用说明： - listener_id，loadbalancer_id，type至少指定一个。 [- 共享型实例的后端服务器组loadbalancer_id和listener_id至少指定一个。 ](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,fcs,dt,hk_tm)
	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`

	// 后端云服务器组的名称。
	Name *string `json:"name,omitempty"`

	// 后端云服务器组所属的项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 后端云服务器组的后端协议。  取值：TCP、UDP、HTTP、HTTPS和QUIC。  使用说明： - listener的protocol为UDP时，pool的protocol必须为UDP或QUIC； - listener的protocol为TCP时pool的protocol必须为TCP； - listener的protocol为HTTP时，pool的protocol必须为HTTP。 - listener的protocol为HTTPS时，pool的protocol必须为HTTP或HTTPS。 - listener的protocol为TERMINATED_HTTPS时，pool的protocol必须为HTTP。 - 若pool的protocol为QUIC，则必须开启session_persistence且type为SOURCE_IP。  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt)  [荷兰region不支持QUIC。](tag:dt,dt_test)
	Protocol string `json:"protocol"`

	SessionPersistence *CreatePoolSessionPersistenceOption `json:"session_persistence,omitempty"`

	SlowStart *CreatePoolSlowStartOption `json:"slow_start,omitempty"`

	// 是否开启删除保护。  取值：false不开启，true开启，默认false。  > 退场时需要先关闭所有资源的删除保护开关。  [不支持该字段，请勿使用。](tag:hws_eu,g42,hk_g42)  [荷兰region不支持该字段，请勿使用。](tag:dt)
	MemberDeletionProtectionEnable *bool `json:"member_deletion_protection_enable,omitempty"`

	// 后端云服务器组关联的虚拟私有云的ID。  使用说明： - 只能挂载到该虚拟私有云下。 - 只能添加该虚拟私有云下的后端服务器或跨VPC的后端服务器。 - type必须指定为instance。  没有指定vpc_id的约束： - 后续添加后端服务器时，vpc_id由后端服务器所在的虚拟私有云确定。
	VpcId *string `json:"vpc_id,omitempty"`

	// 后端服务器组的类型。  取值： - instance：允许任意类型的后端，type指定为该类型时，vpc_id是必选字段。 - ip：只能添加跨VPC后端，type指定为该类型时，vpc_id不允许指定。  使用说明： - 不传表示允许任意类型的后端，并返回type为空字符串。 - listener_id，loadbalancer_id，type至少指定一个。 [- 共享型实例的后端服务器组loadbalancer_id和listener_id至少指定一个。 ](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,fcs,dt,hk_tm)
	Type *string `json:"type,omitempty"`

	// 修改保护状态, 取值： - nonProtection: 不保护，默认值为nonProtection - consoleProtection: 控制台修改保护
	ProtectionStatus *CreatePoolOptionProtectionStatus `json:"protection_status,omitempty"`

	// 设置保护的原因 >仅当protection_status为consoleProtection时有效。
	ProtectionReason *string `json:"protection_reason,omitempty"`

	// 后端是否开启端口透传，开启后，后端服务器端口与前端监听器端口保持一致。  使用说明： - 仅QUIC,TCP,UDP的pool支持。
	AnyPortEnable *bool `json:"any_port_enable,omitempty"`
}

func (o CreatePoolOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePoolOption struct{}"
	}

	return strings.Join([]string{"CreatePoolOption", string(data)}, " ")
}

type CreatePoolOptionProtectionStatus struct {
	value string
}

type CreatePoolOptionProtectionStatusEnum struct {
	NON_PROTECTION     CreatePoolOptionProtectionStatus
	CONSOLE_PROTECTION CreatePoolOptionProtectionStatus
}

func GetCreatePoolOptionProtectionStatusEnum() CreatePoolOptionProtectionStatusEnum {
	return CreatePoolOptionProtectionStatusEnum{
		NON_PROTECTION: CreatePoolOptionProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: CreatePoolOptionProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c CreatePoolOptionProtectionStatus) Value() string {
	return c.value
}

func (c CreatePoolOptionProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePoolOptionProtectionStatus) UnmarshalJSON(b []byte) error {
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
