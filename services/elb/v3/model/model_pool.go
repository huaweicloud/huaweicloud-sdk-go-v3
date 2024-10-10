package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Pool 参数解释：创建后端服务器组请求返回对象。
type Pool struct {

	// 参数解释：后端服务器组的管理状态。  [不支持该字段，请勿使用。](tag:dt,dt_test,hcso_dt)
	AdminStateUp bool `json:"admin_state_up"`

	// 参数解释：后端服务器组的描述信息。
	Description string `json:"description"`

	// 参数解释：后端服务器组关联的健康检查的ID。
	HealthmonitorId string `json:"healthmonitor_id"`

	// 参数解释：后端服务器组的ID。
	Id string `json:"id"`

	// 参数解释：后端服务器组的负载均衡算法。  约束限制： - 当该字段的取值为SOURCE_IP或QUIC_CID时，后端服务器组绑定的后端服务器的weight字段无效。 - 只有pool的protocol为QUIC时，才支持QUIC_CID算法。  取值范围： - ROUND_ROBIN：加权轮询算法。 - LEAST_CONNECTIONS：加权最少连接算法。 - SOURCE_IP：源IP算法。 - QUIC_CID：连接ID算法。 [- 2_TUPLE_HASH：二元组hash算法，仅IP类型的pool支持。 - 3_TUPLE_HASH：三元组hash算法，仅IP类型的pool支持。 - 5_TUPLE_HASH：五元组hash算法，仅IP类型的pool支持。 - IP型pool不指定该字段时，默认设置为5_TUPLE_HASH。](tag:hws_eu)  [不支持QUIC_CID算法。](tag:hws_eu,g42,hk_g42,hcso_dt)  [荷兰region不支持QUIC_CID。](tag:dt,dt_test)
	LbAlgorithm string `json:"lb_algorithm"`

	// 参数解释：后端服务器组关联的监听器ID列表。
	Listeners []ListenerRef `json:"listeners"`

	// 参数解释：后端服务器组关联的负载均衡器ID列表。
	Loadbalancers []LoadBalancerRef `json:"loadbalancers"`

	// 参数解释：后端服务器组中的后端服务器ID列表。
	Members []MemberRef `json:"members"`

	// 参数解释：后端服务器组的名称。
	Name string `json:"name"`

	// 参数解释：后端服务器组所在的项目ID。
	ProjectId string `json:"project_id"`

	// 参数解释：后端服务器组的后端协议。  约束限制： - listener的protocol为UDP时，pool的protocol必须为UDP或QUIC。 - listener的protocol为TCP时pool的protocol必须为TCP。 [- listener的protocol为IP时，pool的protocol必须为IP。](tag:hws_eu) - listener的protocol为HTTP时，pool的protocol必须为HTTP。 - listener的protocol为HTTPS时，pool的protocol必须为HTTP、HTTPS或GRPC。 - listener的protocol为TERMINATED_HTTPS时，pool的protocol必须为HTTP。 - listener的protocol为QUIC时，pool的protocol必须为HTTP、HTTPS或GRPC。 - listener的protocol为TLS时，pool的protocol必须为TLS或TCP。 - 若pool的protocol为QUIC，则必须开启session_persistence且type为SOURCE_IP。 - 若pool的protocol为GRPC，关联监听器必须开启HTTP2。 - 若pool的protocol为TCP,则pool的ip_version字段取值必须是4。  取值范围：TCP、UDP、[IP、](tag:hws_eu)TLS、GRPC、HTTP、HTTPS和QUIC。  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt)  [荷兰region不支持QUIC。](tag:dt,dt_test)
	Protocol string `json:"protocol"`

	SessionPersistence *SessionPersistence `json:"session_persistence"`

	// 参数解释：后端服务器组支持的IP版本。  [取值范围： - 共享型：固定为v4； -  独享型：取值dualstack、v4。当协议为TCP/UDP时，ip_version为dualstack，表示双栈。当协议为HTTP时，ip_version为v4。 ](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs)  [取值范围：dualstack、v4。当协议为TCP/UDP时，ip_version为dualstack，表示双栈。当协议为HTTP时，ip_version为v4。](tag:hcso_dt)  [不支持IPv6，只会返回v4。](tag:dt,dt_test)
	IpVersion string `json:"ip_version"`

	SlowStart *SlowStart `json:"slow_start"`

	// 参数解释：是否开启误删保护。  取值范围：false不开启，true开启。  > 退场时需要先关闭所有资源的删除保护开关。  [不支持该字段，请勿使用。](tag:hws_eu,g42,hk_g42)  [荷兰region不支持该字段，请勿使用。](tag:dt,dt_test)
	MemberDeletionProtectionEnable bool `json:"member_deletion_protection_enable"`

	// 参数解释：创建时间。  取值范围：格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。  [注意：独享型实例的历史数据以及共享型实例下的资源，不返回该字段。 ](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs,dt,hk_tm)
	CreatedAt *string `json:"created_at,omitempty"`

	// 参数解释：更新时间。  取值范围：格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。  [注意：独享型实例的历史数据以及共享型实例下的资源，不返回该字段。 ](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs,dt,hk_tm)
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 参数解释：后端服务器组关联的虚拟私有云的ID。
	VpcId string `json:"vpc_id"`

	// 参数解释：后端服务器组的类型。  取值范围： - instance：允许任意类型的后端，type指定为该类型时，vpc_id是必选字段。 - ip：只能添加跨VPC后端，type指定为该类型时，vpc_id不允许指定。[pool的protocol为IP时，type不允许设置为ip。](tag:hws_eu)] - 空字符串（\"\"）：允许任意类型的后端
	Type string `json:"type"`

	// 参数解释：修改保护状态,。  取值范围： - nonProtection: 不保护。 - consoleProtection: 控制台修改保护。  默认取值：nonProtection
	ProtectionStatus *PoolProtectionStatus `json:"protection_status,omitempty"`

	// 参数解释：设置保护的原因。  参数限制：仅当protection_status为consoleProtection时有效。
	ProtectionReason *string `json:"protection_reason,omitempty"`

	// 参数解释：后端是否开启端口透传。开启后，后端服务器端口与前端监听器端口保持一致。关闭后，请求会转发给后端服务器protocol_port字段指定端口。取值：false不开启，true开启。  约束限制： - 仅QUIC,TCP,UDP的pool支持。
	AnyPortEnable *bool `json:"any_port_enable,omitempty"`

	ConnectionDrain *ConnectionDrain `json:"connection_drain,omitempty"`

	// 参数解释：IP地址组所在的企业项目ID。  [不支持该字段，请勿使用。](tag:dt,dt_test,hcso_dt)
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	PoolHealth *PoolHealth `json:"pool_health,omitempty"`

	// 参数解释：可用区组，如：center
	PublicBorderGroup *string `json:"public_border_group,omitempty"`
}

func (o Pool) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Pool struct{}"
	}

	return strings.Join([]string{"Pool", string(data)}, " ")
}

type PoolProtectionStatus struct {
	value string
}

type PoolProtectionStatusEnum struct {
	NON_PROTECTION     PoolProtectionStatus
	CONSOLE_PROTECTION PoolProtectionStatus
}

func GetPoolProtectionStatusEnum() PoolProtectionStatusEnum {
	return PoolProtectionStatusEnum{
		NON_PROTECTION: PoolProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: PoolProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c PoolProtectionStatus) Value() string {
	return c.value
}

func (c PoolProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoolProtectionStatus) UnmarshalJSON(b []byte) error {
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
