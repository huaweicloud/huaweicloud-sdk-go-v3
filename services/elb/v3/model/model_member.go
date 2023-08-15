package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Member 后端服务器信息。
type Member struct {

	// 后端服务器ID。 >说明： 此处并非ECS服务器的ID，而是ELB为绑定的后端服务器自动生成的member ID。
	Id string `json:"id"`

	// 后端服务器名称。注意：该名称并非ECS名称。
	Name string `json:"name"`

	// 后端服务器所在的项目ID。
	ProjectId string `json:"project_id"`

	// 所在后端服务器组ID。  不支持该字段，请勿使用。
	PoolId *string `json:"pool_id,omitempty"`

	// 后端云服务器的管理状态。  取值：true、false。  虽然创建、更新请求支持该字段，但实际取值决定于后端云服务器对应的弹性云服务器是否存在。若存在，该值为true，否则，该值为false。
	AdminStateUp bool `json:"admin_state_up"`

	// 后端云服务器所在子网的IPv4子网ID或IPv6子网ID。  若所属的LB的跨VPC后端转发特性已开启，则该字段可以不传，表示添加跨VPC的后端服务器。 此时address必须为IPv4地址，所在的pool的协议必须为TCP/HTTP/HTTPS。  使用说明：该子网和关联的负载均衡器的子网必须在同一VPC下。  [不支持IPv6，请勿设置为IPv6子网ID。](tag:dt,dt_test)
	SubnetCidrId *string `json:"subnet_cidr_id,omitempty"`

	// 后端服务器业务端口。 >在开启端口透传的pool下创建member传该字段不生效，可不传该字段。
	ProtocolPort int32 `json:"protocol_port"`

	// 后端云服务器的权重，请求将根据pool配置的负载均衡算法和后端云服务器的权重进行负载分发。 权重值越大，分发的请求越多。权重为0的后端不再接受新的请求。  取值：0-100，默认1。  使用说明：若所在pool的lb_algorithm取值为SOURCE_IP，该字段无效。
	Weight int32 `json:"weight"`

	// 后端服务器对应的IP地址。  使用说明： - 若subnet_cidr_id为空，表示添加跨VPC后端，此时address必须为IPv4地址。 - 若subnet_cidr_id不为空，表示是一个关联到ECS的后端服务器。该IP地址可以是IPv4或IPv6。 但必须在subnet_cidr_id对应的子网网段中。且只能指定为关联ECS的主网卡IP。  [不支持IPv6，请勿设置为IPv6地址。](tag:dt,dt_test)
	Address string `json:"address"`

	// 当前后端服务器的IP地址版本，由后端系统自动根据传入的address字段确定。取值：v4、v6。
	IpVersion string `json:"ip_version"`

	// 设备所有者。  取值： - 空，表示后端服务器未关联到ECS。 - compute:{az_name}，表示关联到ECS，其中{az_name}表示ECS所在可用区名。  不支持该字段，请勿使用。
	DeviceOwner *string `json:"device_owner,omitempty"`

	// 关联的ECS ID，为空表示后端服务器未关联到ECS。  不支持该字段，请勿使用。
	DeviceId *string `json:"device_id,omitempty"`

	// 后端云服务器的健康状态。当status非空时，以status字段中监听器粒度的健康检查状态优先。  取值： - ONLINE：后端云服务器正常。 - NO_MONITOR：后端云服务器所在的服务器组没有健康检查器。 - OFFLINE：后端云服务器关联的ECS服务器不存在或已关机。
	OperatingStatus string `json:"operating_status"`

	// 后端云服务器监听器粒度的的健康状态。 若绑定的监听器在该字段中，则以该字段中监听器对应的operating_stauts为准。 若绑定的监听器不在该字段中，则以外层的operating_status为准。
	Status []MemberStatus `json:"status"`

	// 所属负载均衡器ID。  不支持该字段，请勿使用。
	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`

	// 后端云服务器关联的负载均衡器ID列表。  不支持该字段，请勿使用。
	Loadbalancers *[]ResourceId `json:"loadbalancers,omitempty"`

	// 创建时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。  [注意：独享型实例的历史数据以及共享型实例下的资源，不返回该字段。 ](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,fcs,dt,hk_tm)
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。  [注意：独享型实例的历史数据以及共享型实例下的资源，不返回该字段。 ](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,fcs,dt,hk_tm)
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 后端云服务器的类型。  取值： - ip：跨VPC的member。 - instance：关联到ECS的member。
	MemberType *string `json:"member_type,omitempty"`

	// member关联的实例ID。空表示member关联的实例为非真实设备 （如：跨VPC场景）
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o Member) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Member struct{}"
	}

	return strings.Join([]string{"Member", string(data)}, " ")
}
