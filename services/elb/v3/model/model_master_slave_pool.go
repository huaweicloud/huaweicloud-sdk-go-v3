package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MasterSlavePool struct {

	// **参数解释**：后端服务器组的描述信息。  **取值范围**：不涉及
	Description string `json:"description"`

	// **参数解释**：后端服务器组的ID。  **取值范围**：不涉及
	Id string `json:"id"`

	// **参数解释**：后端服务器组的负载均衡算法。  **取值范围**： - ROUND_ROBIN：加权轮询算法。 - LEAST_CONNECTIONS：加权最少连接算法。 - SOURCE_IP：源IP算法。 - QUIC_CID：连接ID算法。  [不支持QUIC_CID。](tag:tm,hws_eu,g42,hk_g42,hcso_dt)  [荷兰region不支持QUIC_CID。](tag:dt)
	LbAlgorithm string `json:"lb_algorithm"`

	// **参数解释**：后端服务器组关联的监听器ID列表。
	Listeners []ListenerRef `json:"listeners"`

	// **参数解释**：后端服务器组关联的负载均衡器ID列表。
	Loadbalancers []LoadBalancerRef `json:"loadbalancers"`

	// **参数解释**：后端服务器组中的后端服务器列表。
	Members []MasterSlaveMember `json:"members"`

	// **参数解释**：后端服务器组的名称。  **取值范围**：不涉及
	Name string `json:"name"`

	// **参数解释**：后端服务器组所在的项目ID。  **取值范围**：不涉及
	ProjectId string `json:"project_id"`

	// **参数解释**：后端服务器组的后端协议。  **取值范围**：不涉及TCP、UDP、QUIC、TLS。  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt)  [荷兰region不支持QUIC。](tag:dt)
	Protocol string `json:"protocol"`

	SessionPersistence *SessionPersistence `json:"session_persistence"`

	// **参数解释**：后端服务器组支持的IP版本。  [**取值范围**： - 共享型：固定为v4； - 独享型：取值dualstack、v4。当协议为TCP/UDP时，ip_version为dualstack，表示双栈。当协议为HTTP时，ip_version为v4。 ](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,srg,fcs)  [**取值范围**：dualstack、v4。当协议为TCP/UDP时，ip_version为dualstack，表示双栈。当协议为HTTP时，ip_version为v4。](tag:hcso_dt)  [不支持IPv6，只会返回v4。](tag:dt)
	IpVersion string `json:"ip_version"`

	// **参数解释**：创建时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。  **取值范围**：不涉及  [注意：独享型实例的历史数据以及共享型实例下的资源，不返回该字段。](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,srg,fcs,dt,hk_tm)
	CreatedAt string `json:"created_at"`

	// **参数解释**：更新时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。  **取值范围**：不涉及  [注意：独享型实例的历史数据以及共享型实例下的资源，不返回该字段。](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,srg,fcs,dt,hk_tm)
	UpdatedAt string `json:"updated_at"`

	// **参数解释**：后端服务器组关联的虚拟私有云的ID。  **取值范围**：不涉及
	VpcId string `json:"vpc_id"`

	// **参数解释**：后端服务器组的类型。  **取值范围**：不涉及 - instance：允许任意类型的后端，type指定为该类型时，vpc_id是必选字段。 - ip：只能添加IP类型后端，type指定为该类型时，vpc_id不允许指定。 - 空字符串（\"\"）：允许任意类型的后端
	Type string `json:"type"`

	// **参数解释**：资源所属的企业项目ID。  **取值范围**： - \"0\"：表示资源属于default企业项目。 - UUID格式的字符串，表示非默认企业项目。  [不支持该字段，请勿使用。](tag:dt,hcso_dt)
	EnterpriseProjectId string `json:"enterprise_project_id"`

	Healthmonitor *MasterSlaveHealthMonitor `json:"healthmonitor"`

	// **参数解释**：后端是否开启全端口转发。开启后，后端服务器端口与前端监听器端口保持一致。关闭后，请求会转发给后端服务器protocol_port字段指定端口。取值：false 不开启，true 开启。  **取值范围**：不涉及
	AnyPortEnable *bool `json:"any_port_enable,omitempty"`

	ConnectionDrain *ConnectionDrain `json:"connection_drain,omitempty"`

	QuicCidHashStrategy *QuicCidHashStrategy `json:"quic_cid_hash_strategy,omitempty"`
}

func (o MasterSlavePool) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MasterSlavePool struct{}"
	}

	return strings.Join([]string{"MasterSlavePool", string(data)}, " ")
}
