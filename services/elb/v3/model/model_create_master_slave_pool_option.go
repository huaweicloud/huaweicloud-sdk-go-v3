package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMasterSlavePoolOption 创建主备主机组请求
type CreateMasterSlavePoolOption struct {

	// 后端服务器组的描述信息。
	Description *string `json:"description,omitempty"`

	// 后端服务器组的负载均衡算法。  取值： - ROUND_ROBIN：加权轮询算法。 - LEAST_CONNECTIONS：加权最少连接算法。 - SOURCE_IP：源IP算法。 - QUIC_CID：连接ID算法。  [不支持QUIC_CID。](tag:tm,hws_eu,g42,hk_g42,hcso_dt)  [荷兰region不支持QUIC_CID。](tag:dt,dt_test)
	LbAlgorithm string `json:"lb_algorithm"`

	// 后端服务器组关联的LB的ID。  使用说明：listener_id，loadbalancer_id，type至少指定一个。
	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`

	// 后端服务器组关联的监听器的ID。  使用说明：listener_id，loadbalancer_id，type至少指定一个。
	ListenerId *string `json:"listener_id,omitempty"`

	// 后端服务器组的名称。
	Name *string `json:"name,omitempty"`

	// 后端服务器组所属的项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 后端服务器组的后端协议。  取值：TCP、UDP、QUIC、TLS。  使用说明： - listener的protocol为UDP时，pool的protocol必须为UDP或QUIC。 - listener的protocol为TCP时，pool的protocol必须为TCP。 - listener的protocol为TLS时，pool的protocol必须为TLS或TCP（且只能使用ip_version为v4的TCP pool）。 - 其他协议监听器不支持主备后端服务器组。  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt)  [荷兰region不支持QUIC。](tag:dt,dt_test)
	Protocol string `json:"protocol"`

	SessionPersistence *CreatePoolSessionPersistenceOption `json:"session_persistence,omitempty"`

	// 后端服务器组关联的虚拟私有云的ID。  指定了vpc_id的约束： - 只能挂载到该虚拟私有云下。 - 只能添加该虚拟私有云下的后端服务器或跨VPC的后端服务器。 - type必须指定为instance。  没有指定vpc_id的约束： - 后续添加后端服务器时，vpc_id由后端服务器所在的虚拟私有云确定。
	VpcId *string `json:"vpc_id,omitempty"`

	// 后端服务器组的类型。  取值： - instance：允许任意类型的后端，type指定为该类型时，vpc_id是必选字段。 - ip：只能添加跨VPC后端，type指定为该类型时，vpc_id不允许指定。
	Type string `json:"type"`

	// 后端服务器组支持的IP版本。  [取值： - 共享型：固定为v4； -  独享型：取值dualstack、v4、v6。当协议为TCP/UDP时，ip_version为dualstack，表示双栈。  当协议为HTTP时，ip_version为v4。 ](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,fcs)  [取值：dualstack、v4、v6。当协议为TCP/UDP时，ip_version为dualstack，表示双栈。 当协议为HTTP时，ip_version为v4。](tag:hcso_dt)  [不支持IPv6，只会返回v4。](tag:dt,dt_test)
	IpVersion *string `json:"ip_version,omitempty"`

	// 主备主机组的后端服务器。 只能添加2个后端服务器，必须有一个为主，一个为备。
	Members []CreateMasterSlaveMemberOption `json:"members"`

	Healthmonitor *CreateMasterSlaveHealthMonitorOption `json:"healthmonitor"`

	// 后端是否开启端口透传。开启后，后端服务器端口与前端监听器端口保持一致。关闭后，请求会转发给后端服务器protocol_port字段指定端口。取值：false不开启，true开启。  使用说明： - 仅QUIC,TCP,UDP的pool支持。
	AnyPortEnable *bool `json:"any_port_enable,omitempty"`

	ConnectionDrain *ConnectionDrain `json:"connection_drain,omitempty"`
}

func (o CreateMasterSlavePoolOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMasterSlavePoolOption struct{}"
	}

	return strings.Join([]string{"CreateMasterSlavePoolOption", string(data)}, " ")
}
