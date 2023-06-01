package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 创建云服务器组请求返回对象
type Pool struct {

	// 后端云服务器组的管理状态，只支持设置为true。  不支持该字段，请勿使用。
	AdminStateUp bool `json:"admin_state_up"`

	// 后端云服务器组的描述信息。
	Description string `json:"description"`

	// 后端云服务器组关联的健康检查的ID。
	HealthmonitorId string `json:"healthmonitor_id"`

	// 后端云服务器组的ID。
	Id string `json:"id"`

	// 后端云服务器组的负载均衡算法。  取值： - ROUND_ROBIN：加权轮询算法。 - LEAST_CONNECTIONS：加权最少连接算法。 - SOURCE_IP：源IP算法。 - QUIC_CID：连接ID算法。  使用说明： - 当该字段的取值为SOURCE_IP时，后端云服务器组绑定的后端云服务器的weight字段无效。 - 只有pool的protocol为QUIC时，才支持QUIC_CID算法。  [不支持QUIC_CID算法。](tag:hws_eu,g42,hk_g42,hcso_dt)  [荷兰region不支持QUIC_CID。](tag:dt,dt_test)
	LbAlgorithm string `json:"lb_algorithm"`

	// 后端云服务器组关联的监听器ID列表。
	Listeners []ListenerRef `json:"listeners"`

	// 后端云服务器组关联的负载均衡器ID列表。
	Loadbalancers []LoadBalancerRef `json:"loadbalancers"`

	// 后端云服务器组中的后端云服务器ID列表。
	Members []MemberRef `json:"members"`

	// 后端云服务器组的名称。
	Name string `json:"name"`

	// 后端云服务器组所在的项目ID。
	ProjectId string `json:"project_id"`

	// 后端云服务器组的后端协议。  取值：TCP、UDP、HTTP、HTTPS、QUIC。  使用说明： - listener的protocol为UDP时，pool的protocol必须为UDP或QUIC； - listener的protocol为TCP时pool的protocol必须为TCP； - listener的protocol为HTTP时，pool的protocol必须为HTTP。 - listener的protocol为HTTPS时，pool的protocol必须为HTTP或HTTPS。 - listener的protocol为TERMINATED_HTTPS时，pool的protocol必须为HTTP。 - 若pool的protocol为QUIC，则必须开启session_persistence且type为SOURCE_IP。  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt)  [荷兰region不支持QUIC。](tag:dt,dt_test)
	Protocol string `json:"protocol"`

	SessionPersistence *SessionPersistence `json:"session_persistence"`

	// 后端云服务器组支持的IP版本。  [取值： - 共享型：固定为v4； -  独享型：取值dualstack、v4、v6。当协议为TCP/UDP时，ip_version为dualstack，表示双栈。  当协议为HTTP时，ip_version为v4。 ](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,fcs)  [取值：dualstack、v4、v6。当协议为TCP/UDP时，ip_version为dualstack，表示双栈。 当协议为HTTP时，ip_version为v4。](tag:hcso_dt)  [不支持IPv6，只会返回v4。](tag:dt,dt_test)
	IpVersion *string `json:"ip_version,omitempty"`

	SlowStart *SlowStart `json:"slow_start"`

	// 是否开启误删保护。  取值：false不开启，true开启。  > 退场时需要先关闭所有资源的删除保护开关。  [不支持该字段，请勿使用。](tag:hws_eu,g42,hk_g42)  [荷兰region不支持该字段，请勿使用。](tag:dt)
	MemberDeletionProtectionEnable bool `json:"member_deletion_protection_enable"`

	// 创建时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。  [注意：独享型实例的历史数据以及共享型实例下的资源，不返回该字段。 ](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,fcs,dt,hk_tm)
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。  [注意：独享型实例的历史数据以及共享型实例下的资源，不返回该字段。 ](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,fcs,dt,hk_tm)
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 后端云服务器组关联的虚拟私有云的ID。
	VpcId string `json:"vpc_id"`

	// 后端服务器组的类型。  取值： - instance：允许任意类型的后端，type指定为该类型时，vpc_id是必选字段。 - ip：只能添加跨VPC后端，type指定为该类型时，vpc_id不允许指定。 - 空字符串（\"\"）：允许任意类型的后端
	Type string `json:"type"`

	// 修改保护状态, 取值： - nonProtection: 不保护，默认值为nonProtection - consoleProtection: 控制台修改保护
	ProtectionStatus *PoolProtectionStatus `json:"protection_status,omitempty"`

	// 设置保护的原因 >仅当protection_status为consoleProtection时有效。
	ProtectionReason *string `json:"protection_reason,omitempty"`
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
