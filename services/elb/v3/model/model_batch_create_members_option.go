package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateMembersOption 批量添加member请求参数。
type BatchCreateMembersOption struct {

	// 后端服务器名称。
	Name *string `json:"name,omitempty"`

	// 参数解释：后端服务器的可用区。 约束限制：  仅支持iptarget类型的后端服务器设置该字段。且后端服务器组开启可用区亲和时，iptarget类型的后端服务器必须配置该字段为有效非空值。 取值范围：本region中ECS可选择的可用区。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 后端服务器的对应的IP地址，这个IP必须在subnet_cidr_id字段的子网网段中。例如：192.168.3.11。  subnet_cidr_id为空代表添加跨VPC后端，此时address必须为**私网IPv4**地址。
	Address string `json:"address"`

	// 后端服务器业务端口。  [网关型LB，即pool协议为IP时，protocol_port必须设置为0。](tag:hws_eu)  >在开启端口透传的pool下创建member传该字段不生效，可不传该字段。
	ProtocolPort int32 `json:"protocol_port"`

	// 后端服务器所在的子网，可以是IPv4或IPv6子网。若是IPv4子网，使用对应子网的子网ID（neutron_subnet_id）；若是IPv6子网，使用对应子网的网络ID（neutron_network_id）。  ipv4子网的子网ID可以通过GET https://{VPC_Endpoint}/v1/{project_id}/subnets 响应参数中的neutron_subnet_id得到  ipv6子网的网络ID可以通过GET https://{VPC_Endpoint}/v1/{project_id}/subnets 响应参数中的neutron_network_id得到  使用说明： - 该子网和关联的负载均衡器的子网必须在同一VPC下。 - 若所属LB的跨VPC后端转发已开启（ip_target_enable=true），则该字段可以不传，表示添加跨VPC的后端服务器。此时address必须为**私网IPv4**地址，所在的pool的协议必须为UDP/TCP/TLS/HTTP/HTTPS/QUIC/GRPC。 - 若所属LB未开启跨VPC后端转发，该参数必填。 [- 网关型LB，即pool协议为IP时，必须指定该子网，且必须和负载均衡器的子网在同一个VPC下，但不能相同。](tag:hws_eu)  [不支持IPv6，请勿设置为IPv6子网ID。](tag:dt)
	SubnetCidrId *string `json:"subnet_cidr_id,omitempty"`

	// 后端服务器的权重，请求将根据pool配置的负载均衡算法和后端服务器的权重进行负载分发。 权重值越大，分发的请求越多。权重为0的后端不再接受新的请求。  取值：0-100，默认1。  使用说明：若所在pool的lb_algorithm取值为SOURCE_IP或QUIC_CID，该字段无效。
	Weight *int32 `json:"weight,omitempty"`
}

func (o BatchCreateMembersOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateMembersOption struct{}"
	}

	return strings.Join([]string{"BatchCreateMembersOption", string(data)}, " ")
}
