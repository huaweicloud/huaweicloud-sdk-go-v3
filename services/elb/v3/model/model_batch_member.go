package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量创建后端服务器响应参数
type BatchMember struct {
	// 后端服务器ID。

	Id string `json:"id"`
	// 后端服务器名称。

	Name string `json:"name"`
	// 后端服务器所在的项目ID。

	ProjectId string `json:"project_id"`
	// 后端云服务器的管理状态。取值：true、false。  虽然创建、更新请求支持该字段，但实际取值决定于后端云服务器对应的弹性云服务器是否存在。若存在，该值为true，否则，该值为false。

	AdminStateUp bool `json:"admin_state_up"`
	// 后端云服务器所在子网的IPv4子网ID或IPv6子网ID。  若所属的LB的跨VPC后端转发特性已开启，则该字段可以不传，表示添加跨VPC的后端服务器。此时address必须为IPv4地址，所在的pool的协议必须为TCP/HTTP/HTTPS。  使用说明： - 该子网和关联的负载均衡器的子网必须在同一VPC下。 [不支持IPv6，请勿设置为IPv6子网ID。](tag:dt,dt_test)

	SubnetCidrId *string `json:"subnet_cidr_id,omitempty"`
	// 后端服务器业务端口号。

	ProtocolPort int32 `json:"protocol_port"`
	// 后端云服务器的权重，请求将根据pool配置的负载均衡算法和后端云服务器的权重进行负载分发。权重值越大，分发的请求越多。权重为0的后端不再接受新的请求。 取值：0-100，默认1。 使用说明：  - 若所在pool的lb_algorithm取值为SOURCE_IP，该字段无效。

	Weight int32 `json:"weight"`
	// 后端服务器对应的IP地址。 使用说明：  - 若subnet_cidr_id为空，表示添加跨VPC后端，此时address必须为IPv4地址。  - 若subnet_cidr_id不为空，表示是一个关联到ECS的后端服务器。该IP地址可以是IPv4或IPv6。但必须在subnet_cidr_id对应的子网网段中。且只能指定为关联ECS的主网卡IP。 [不支持IPv6，请勿设置为IPv6地址。](tag:dt,dt_test)

	Address string `json:"address"`
	// 后端云服务器的健康状态。取值： - ONLINE：后端云服务器正常。 - NO_MONITOR：后端云服务器所在的服务器组没有健康检查器。 - OFFLINE：后端云服务器关联的ECS服务器不存在或已关机。

	OperatingStatus string `json:"operating_status"`
	// 后端云服务器的类型。取值： - ip：跨VPC的member。 - instance：关联到ECS的member。

	MemberType *string `json:"member_type,omitempty"`
	// member关联的实例ID，空表示跨VPC场景的member。

	InstanceId *string `json:"instance_id,omitempty"`
	// IP地址对应的VPC port ID

	PortId string `json:"port_id"`
	// 当前后端服务器创建结果状态。取值： - successful：添加成功。 - existed：member已存在。

	RetStatus string `json:"ret_status"`
}

func (o BatchMember) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchMember struct{}"
	}

	return strings.Join([]string{"BatchMember", string(data)}, " ")
}
