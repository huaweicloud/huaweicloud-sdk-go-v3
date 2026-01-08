package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Member struct {

	// **参数解释**：后端服务器ID。 **取值范围**：不涉及 > 此处并非ECS服务器的ID，而是ELB为绑定的后端服务器自动生成的member ID。
	Id string `json:"id"`

	// **参数解释**：后端服务器所在的可用区。  **取值范围**：不涉及
	AvailabilityZone string `json:"availability_zone"`

	// **参数解释**：后端服务器名称。  **取值范围**：不涉及  注意：该名称并非ECS名称。
	Name string `json:"name"`

	// **参数解释**：后端服务器所在的项目ID。  **取值范围**：不涉及
	ProjectId string `json:"project_id"`

	// **参数解释**：后端服务器的管理状态。  **取值范围**：true、false。
	AdminStateUp bool `json:"admin_state_up"`

	// **参数解释**：后端服务器所在的子网，可以是IPv4或IPv6子网。若是IPv4子网，使用对应子网的子网ID（neutron_subnet_id）；若是IPv6子网，使用对应子网的网络ID（neutron_network_id）。 ipv4子网的子网ID可以通过GET https://{VPC_Endpoint}/v1/{project_id}/subnets 响应参数中的neutron_subnet_id得到 ipv6子网的网络ID可以通过GET https://{VPC_Endpoint}/v1/{project_id}/subnets 响应参数中的neutron_network_id得到  **取值范围**：不涉及  [不支持IPv6，请勿设置为IPv6子网ID。](tag:dt)
	SubnetCidrId *string `json:"subnet_cidr_id,omitempty"`

	// **参数解释**：后端服务器业务端口。  **取值范围**：不涉及  >在开启端口透传的pool下创建member传该字段不生效，可不传该字段。
	ProtocolPort int32 `json:"protocol_port"`

	// **参数解释**：后端服务器的权重，请求将根据pool配置的负载均衡算法和后端服务器的权重进行负载分发。 权重值越大，分发的请求越多。权重为0的后端不再接受新的请求。  **取值范围**：0-100
	Weight int32 `json:"weight"`

	// **参数解释**：后端服务器对应的IP地址。  **取值范围**：不涉及  [不支持IPv6，请勿设置为IPv6地址。](tag:dt)
	Address string `json:"address"`

	// **参数解释**：当前后端服务器的IP地址版本，由后端系统自动根据传入的address字段确定。  **取值范围**：v4、v6
	IpVersion string `json:"ip_version"`

	// **参数解释**：后端服务器的健康状态。当status非空时，以status字段中监听器粒度的健康检查状态优先。  **取值范围**： - NO_MONITOR：后端服务器所在的服务器组没有开启健康检查。 - INITIAL：初始化中，表示负载均衡实例配置了健康检查，但查不到数据。 - ONLINE：后端服务器正常。 - OFFLINE：后端服务器关联的ECS服务器不存在或已关机。 - UNKNOWN：未关联LB实例的pool下的member，或者创建后从未关联ECS的云服务器类型member，状态置为UNKNOWN。
	OperatingStatus string `json:"operating_status"`

	// **参数解释**：后端服务器监听器粒度的的健康状态。  **取值范围**：不涉及
	Status []MemberStatus `json:"status"`

	Reason *MemberHealthCheckFailedReason `json:"reason,omitempty"`

	// **参数解释**：创建时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。  **取值范围**：不涉及  [注意：独享型实例的历史数据以及共享型实例下的资源，不返回该字段。](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,srg,fcs,dt,hk_tm)
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释**：更新时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。  **取值范围**：不涉及  [注意：独享型实例的历史数据以及共享型实例下的资源，不返回该字段。](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,srg,fcs,dt,hk_tm)
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释**：后端服务器的类型。  **取值范围**： - ip：IP类型的member。 - instance：关联到ECS的member。
	MemberType *string `json:"member_type,omitempty"`

	// **参数解释**：member关联的实例ID。空表示member关联的实例为非真实设备 （如：跨VPC场景）  **取值范围**：不涉及
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o Member) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Member struct{}"
	}

	return strings.Join([]string{"Member", string(data)}, " ")
}
