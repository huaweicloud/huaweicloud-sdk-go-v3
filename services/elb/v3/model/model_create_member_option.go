package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMemberOption 创建后端服务器请求参数
type CreateMemberOption struct {

	// 后端服务器对应的IP地址。  使用说明： - 若subnet_cidr_id为空，表示添加跨VPC后端，此时address必须为IPv4地址。 - 若subnet_cidr_id不为空，表示是一个关联到ECS的后端服务器。该IP地址可以是私网IPv4或IPv6。 但必须在subnet_cidr_id对应的子网网段中。且只能指定为关联ECS的主网卡的内网IP。  [ 不支持IPv6，请勿设置为IPv6地址。](tag:dt,dt_test)
	Address string `json:"address"`

	// 后端云服务器的管理状态。  取值：true、false。  虽然创建、更新请求支持该字段，但实际取值决定于后端云服务器对应的弹性云服务器是否存在。若存在，该值为true，否则，该值为false。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 后端云服务器名称。注意：该名称并非ECS名称，若不传则返回为空。
	Name *string `json:"name,omitempty"`

	// 后端云服务器所在的项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 后端服务器业务端口。 >在开启端口透传的pool下创建member传该字段不生效，可不传该字段。
	ProtocolPort *int32 `json:"protocol_port,omitempty"`

	// 后端云服务器所在的子网，可以是IPv4或IPv6子网。若是IPv4子网，使用对应子网的子网ID（neutron_subnet_id）；若是IPv6子网，使用对应子网的网络ID（neutron_network_id）。  ipv4子网的子网ID可以通过GET https://{VPC_Endpoint}/v1/{project_id}/subnets 响应参数中的neutron_subnet_id得到  ipv6子网的网络ID可以通过GET https://{VPC_Endpoint}/v1/{project_id}/subnets 响应参数中的neutron_network_id得到  使用说明： - 该子网和关联的负载均衡器的子网必须在同一VPC下。 - 若所属LB的跨VPC后端转发已开启（ip_target_enable=true），则该字段可以不传，表示添加跨VPC的后端服务器。 此时address必须为IPv4地址，所在的pool的协议必须为TCP/HTTP/HTTPS。 - 若所属LB未开启跨VPC后端转发，该参数必填。  [不支持IPv6，请勿设置为IPv6子网ID。](tag:dt,dt_test)
	SubnetCidrId *string `json:"subnet_cidr_id,omitempty"`

	// 后端云服务器的权重，请求将根据pool配置的负载均衡算法和后端云服务器的权重进行负载分发。 权重值越大，分发的请求越多。权重为0的后端不再接受新的请求。  取值：0-100，默认1。  使用说明：若所在pool的lb_algorithm取值为SOURCE_IP，该字段无效。
	Weight *int32 `json:"weight,omitempty"`
}

func (o CreateMemberOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMemberOption struct{}"
	}

	return strings.Join([]string{"CreateMemberOption", string(data)}, " ")
}
