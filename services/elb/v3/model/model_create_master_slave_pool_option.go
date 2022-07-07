package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建主备主机组请求
type CreateMasterSlavePoolOption struct {

	// 后端云服务器组的描述信息。
	Description *string `json:"description,omitempty"`

	// 后端云服务器组的负载均衡算法。  取值： - ROUND_ROBIN：加权轮询算法。 - LEAST_CONNECTIONS：加权最少连接算法。 - SOURCE_IP：源IP算法。 - QUIC_CID：连接ID算法。
	LbAlgorithm string `json:"lb_algorithm"`

	// 后端云服务器组关联的LB的ID。  使用说明：listener_id，loadbalancer_id，type至少指定一个。
	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`

	// 后端云服务器组关联的监听器的ID。  使用说明：listener_id，loadbalancer_id，type至少指定一个。
	ListenerId *string `json:"listener_id,omitempty"`

	// 后端云服务器组的名称。
	Name *string `json:"name,omitempty"`

	// 后端云服务器组所属的项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 后端云服务器组的后端协议。  取值：TCP、UDP、QUIC。  使用说明： - listener的protocol为UDP时，pool的protocol必须为UDP或QUIC。 - listener的protocol为TCP时pool的protocol必须为TCP。
	Protocol string `json:"protocol"`

	SessionPersistence *CreatePoolSessionPersistenceOption `json:"session_persistence,omitempty"`

	// 后端云服务器组关联的虚拟私有云的ID。   指定了vpc_id的约束： - 只能挂载到该虚拟私有云下。 - 只能添加该虚拟私有云下的后端服务器或跨VPC的后端服务器。 - type必须指定为instance。   没有指定vpc_id的约束： - 后续添加后端服务器时，vpc_id由后端服务器所在的虚拟私有云确定。
	VpcId *string `json:"vpc_id,omitempty"`

	// 后端服务器组的类型。   取值：  - instance：允许任意类型的后端，type指定为该类型时，vpc_id是必选字段。  - ip：只能添加跨VPC后端，type指定为该类型时，vpc_id不允许指定。
	Type string `json:"type"`

	// 主备主机组的后端服务器。 只能添加2个后端云服务器，必须有一个为主，一个为备。
	Members []CreateMasterSlaveMemberOption `json:"members"`

	Healthmonitor *CreateMasterSlaveHealthMonitorOption `json:"healthmonitor"`
}

func (o CreateMasterSlavePoolOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMasterSlavePoolOption struct{}"
	}

	return strings.Join([]string{"CreateMasterSlavePoolOption", string(data)}, " ")
}
