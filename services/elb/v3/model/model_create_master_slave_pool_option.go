package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMasterSlavePoolOption **参数解释**：创建主备服务器组请求参数。  **约束限制**：不涉及
type CreateMasterSlavePoolOption struct {

	// **参数解释**：后端服务器组的描述信息。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**：后端服务器组的负载均衡算法。  **约束限制**：不涉及  **取值范围**： - ROUND_ROBIN：加权轮询算法。 - LEAST_CONNECTIONS：加权最少连接算法。 - SOURCE_IP：源IP算法。 - QUIC_CID：连接ID算法。  **默认取值**：不涉及  [不支持QUIC_CID。](tag:tm,hws_eu,g42,hk_g42,hcso_dt)  [荷兰region不支持QUIC_CID。](tag:dt)
	LbAlgorithm string `json:"lb_algorithm"`

	// **参数解释**：后端服务器组关联的LB的ID。  **约束限制**：listener_id，loadbalancer_id，type至少指定一个。  **取值范围**：不涉及  **默认取值**：不涉及
	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`

	// **参数解释**：后端服务器组关联的监听器的ID。  **约束限制**：listener_id，loadbalancer_id，type至少指定一个。  **取值范围**：不涉及  **默认取值**：不涉及
	ListenerId *string `json:"listener_id,omitempty"`

	// **参数解释**：后端服务器组的名称。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**：项目ID。获取方式请参见[获取项目ID](elb_fl_0008.xml)。  **约束限制**：不涉及  **取值范围**：长度为32个字符，由小写字母和数字组成。  **默认取值**：不涉及  > 该字段实际无效，最终使用url中的project_id。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**：后端服务器组的后端协议。  **约束限制**： - listener的protocol为UDP时，pool的protocol必须为UDP或QUIC。 - listener的protocol为TCP时，pool的protocol必须为TCP。 - listener的protocol为TLS时，pool的protocol必须为TLS或TCP（且只能使用ip_version为v4的TCP pool）。 - 其他协议监听器不支持主备后端服务器组。  **取值范围**：TCP、UDP、QUIC、TLS。  **默认取值**：不涉及  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt)  [荷兰region不支持QUIC。](tag:dt)
	Protocol string `json:"protocol"`

	SessionPersistence *CreatePoolSessionPersistenceOption `json:"session_persistence,omitempty"`

	// **参数解释**：后端服务器组关联的虚拟私有云的ID。  **约束限制**： - 只能挂载到该虚拟私有云下。 - 只能添加该虚拟私有云下的后端服务器或跨VPC的后端服务器。 - type必须指定为instance。 [- pool的protocol为IP时，必须指定vpc_id，且与LB的vpc_id相同。](tag:hws_eu) - 若未指定vpc_id，则后续添加后端服务器时，vpc_id由后端服务器所在的虚拟私有云确定。  **取值范围**：不涉及  **默认取值**：不涉及
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释**：后端服务器组的类型。  **约束限制**：不涉及  **取值范围**： - instance：允许任意类型的后端，type指定为该类型时，vpc_id是必选字段。 - ip：只能添加IP类型后端，type指定为该类型时，vpc_id不允许指定。  **默认取值**：不涉及
	Type string `json:"type"`

	// **参数解释**：后端服务器组支持的IP版本。  **约束限制**：不涉及  [**取值范围**： - 共享型：固定为v4； - 独享型：取值dualstack、v4。当协议为TCP/UDP时，ip_version为dualstack，表示双栈。当协议为HTTP时，ip_version为v4。 ](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,srg,fcs)  [**取值范围**：dualstack、v4。当协议为TCP/UDP时，ip_version为dualstack，表示双栈。当协议为HTTP时，ip_version为v4。](tag:hcso_dt)  **默认取值**：不涉及  [不支持IPv6，只会返回v4。](tag:dt)
	IpVersion *string `json:"ip_version,omitempty"`

	// **参数解释**：主备服务器组的后端服务器。  **约束限制**：只能添加2个后端服务器，必须有一个为主，一个为备。
	Members []CreateMasterSlaveMemberOption `json:"members"`

	Healthmonitor *CreateMasterSlaveHealthMonitorOption `json:"healthmonitor"`

	// **参数解释**：后端是否开启全端口转发。开启后，后端服务器端口与前端监听器端口保持一致。关闭后，请求会转发给后端服务器protocol_port字段指定端口。  **约束限制**：仅QUIC,TCP,UDP的pool支持。  **取值范围**：false 不开启，true 开启。  **默认取值**：不涉及
	AnyPortEnable *bool `json:"any_port_enable,omitempty"`

	ConnectionDrain *ConnectionDrain `json:"connection_drain,omitempty"`

	QuicCidHashStrategy *QuicCidHashStrategy `json:"quic_cid_hash_strategy,omitempty"`
}

func (o CreateMasterSlavePoolOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMasterSlavePoolOption struct{}"
	}

	return strings.Join([]string{"CreateMasterSlavePoolOption", string(data)}, " ")
}
