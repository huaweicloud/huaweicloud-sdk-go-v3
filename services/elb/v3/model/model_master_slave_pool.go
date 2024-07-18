package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MasterSlavePool 创建云服务器组请求返回对象
type MasterSlavePool struct {

	// 后端云服务器组的描述信息。
	Description string `json:"description"`

	// 后端云服务器组的ID。
	Id string `json:"id"`

	// 后端云服务器组的负载均衡算法。  取值： - ROUND_ROBIN：加权轮询算法。 - LEAST_CONNECTIONS：加权最少连接算法。 - SOURCE_IP：源IP算法。 - QUIC_CID：连接ID算法。  使用说明： - 当该字段的取值为SOURCE_IP时，后端云服务器组绑定的后端云服务器的weight字段无效。 - 只有pool的protocol为QUIC时，才支持QUIC_CID算法。  [不支持QUIC_CID。](tag:tm,hws_eu,g42,hk_g42,hcso_dt)  [荷兰region不支持QUIC_CID。](tag:dt,dt_test)
	LbAlgorithm string `json:"lb_algorithm"`

	// 后端云服务器组关联的监听器ID列表。
	Listeners []ListenerRef `json:"listeners"`

	// 后端云服务器组关联的负载均衡器ID列表。
	Loadbalancers []LoadBalancerRef `json:"loadbalancers"`

	// 后端云服务器组中的后端云服务器列表。
	Members []MasterSlaveMember `json:"members"`

	// 后端云服务器组的名称。
	Name string `json:"name"`

	// 后端云服务器组所在的项目ID。
	ProjectId string `json:"project_id"`

	// 后端云服务器组的后端协议。  取值：TCP、UDP、QUIC、TLS。  使用说明： - listener的protocol为UDP时，pool的protocol必须为UDP或QUIC。 - listener的protocol为TCP时，pool的protocol必须为TCP。 - listener的protocol为TLS时，pool的protocol必须为TLS或TCP。 - 其他协议监听器不支持主备后端服务器组。  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt)  [荷兰region不支持QUIC。](tag:dt,dt_test)
	Protocol string `json:"protocol"`

	SessionPersistence *SessionPersistence `json:"session_persistence"`

	// 后端云服务器组支持的IP版本。  [取值： - 共享型：固定为v4； -  独享型：取值dualstack、v4、v6。当协议为TCP/UDP时，ip_version为dualstack，表示双栈。  当协议为HTTP时，ip_version为v4。 ](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs)  [取值：dualstack、v4、v6。当协议为TCP/UDP时，ip_version为dualstack，表示双栈。 当协议为HTTP时，ip_version为v4。](tag:hcso_dt)  [不支持IPv6，只会返回v4。](tag:dt,dt_test)
	IpVersion string `json:"ip_version"`

	// 创建时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。  [注意：独享型实例的历史数据以及共享型实例下的资源，不返回该字段。 ](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs,dt,hk_tm)
	CreatedAt string `json:"created_at"`

	// 更新时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。  [注意：独享型实例的历史数据以及共享型实例下的资源，不返回该字段。 ](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs,dt,hk_tm)
	UpdatedAt string `json:"updated_at"`

	// 后端云服务器组关联的虚拟私有云的ID。
	VpcId string `json:"vpc_id"`

	// 后端服务器组的类型。  取值： - instance：允许任意类型的后端，type指定为该类型时，vpc_id是必选字段。 - ip：只能添加跨VPC后端，type指定为该类型时，vpc_id不允许指定。 - 空字符串（\"\"）：允许任意类型的后端
	Type string `json:"type"`

	// 后端服务器组的企业项目ID。无论创建什么企业项目，都在默认企业项目下。  [不支持该字段，请勿使用。](tag:dt,dt_test,hcso_dt)
	EnterpriseProjectId string `json:"enterprise_project_id"`

	Healthmonitor *MasterSlaveHealthMonitor `json:"healthmonitor"`

	// 后端是否开启端口透传。开启后，后端服务器端口与前端监听器端口保持一致。关闭后，请求会转发给后端服务器protocol_port字段指定端口。取值：false不开启，true开启。  使用说明： - 仅QUIC,TCP,UDP的pool支持。
	AnyPortEnable *bool `json:"any_port_enable,omitempty"`

	ConnectionDrain *ConnectionDrain `json:"connection_drain,omitempty"`
}

func (o MasterSlavePool) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MasterSlavePool struct{}"
	}

	return strings.Join([]string{"MasterSlavePool", string(data)}, " ")
}
